package graph

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/vikstrous/dataloadgen"

	"nordkapp42/graph/model"
)

func ForLoaders(ctx context.Context) *Loaders {
	return ctx.Value("loaders").(*Loaders)
}

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	MessageLoader *dataloadgen.Loader[int, *model.Message]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders() *Loaders {
	// define the data loader
	return &Loaders{
		MessageLoader: dataloadgen.NewLoader(getModels[int, model.Message], dataloadgen.WithWait(time.Millisecond)),
	}
}

func getModels[T comparable, M model.Model[T]](ctx context.Context, keys []T) ([]*M, []error) {
	var dbModels []*M
	// Using Bun, we adjust the query method to fit Bun's API.
	// Bun's handling of `WhereIn` is similar to go-pg, but ensure you're using the correct
	// query method calls for Bun.
	db := ForDB(ctx)
	err := db.NewSelect().Model(&dbModels).Where("id IN (?)", bun.In(keys)).Scan(ctx)
	if err != nil {
		return []*M{}, []error{err}
	}
	// Mapping message IDs to messages for quick lookup.
	m := make(map[T]*M)
	for _, dbModel := range dbModels {
		m[M(*dbModel).GetID()] = dbModel
	}
	// Reassembling the results in the order of keys.
	models := make([]*M, len(keys))
	errs := make([]error, len(keys))
	for i, key := range keys {
		if model, ok := m[key]; ok {
			models[i] = model
		} else {
			// Handle the case where a key does not have a corresponding user.
			errs = append(errs, fmt.Errorf("no element found for key: %v", key))
			models[i] = nil // Keep place with nil if element is not found.
		}
	}
	return models, errs
}
