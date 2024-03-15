package loaders

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/vikstrous/dataloadgen"

	"nordkapp42/graph/model"
	"nordkapp42/utils"
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
		// MessageLoader: dataloadgen.NewLoader(getModels[int, model.Message], dataloadgen.WithWait(time.Millisecond)),
		MessageLoader: dataloadgen.NewLoader(getMessages, dataloadgen.WithWait(time.Millisecond)),
	}
}

// type Model[T comparable] struct {
// 	ID T
// }

// func getModels[T comparable, M any](ctx context.Context, keys []T) ([]*M, []error) {
// 	var dbModels []*M
// 	// Using Bun, we adjust the query method to fit Bun's API.
// 	// Bun's handling of `WhereIn` is similar to go-pg, but ensure you're using the correct
// 	// query method calls for Bun.
// 	db := utils.ForDB(ctx)
// 	err := db.NewSelect().Model(&dbModels).Where("id IN (?)", bun.In(keys)).Scan(ctx)
// 	if err != nil {
// 		return []*M{}, []error{err}
// 	}
// 	// Mapping message IDs to messages for quick lookup.
// 	m := make(map[T]*M)
// 	for _, dbModel := range dbModels {
// 		var dummy any
// 		dummy = *dbModel
// 		model := dummy.(Model[T])
// 		m[model.ID] = dbModel
// 	}
// 	// Reassembling the results in the order of keys.
// 	models := make([]*M, len(keys))
// 	errs := make([]error, len(keys))
// 	for i, key := range keys {
// 		if model, ok := m[key]; ok {
// 			models[i] = model
// 		} else {
// 			// Handle the case where a key does not have a corresponding user.
// 			errs = append(errs, fmt.Errorf("no element found for key: %d", key))
// 			models[i] = nil // Keep place with nil if element is not found.
// 		}
// 	}
// 	// return elements, errs
// 	return models, errs
// }

func getMessages(ctx context.Context, keys []int) ([]*model.Message, []error) {
	var dbMessages []*model.Message
	// Using Bun, we adjust the query method to fit Bun's API.
	// Bun's handling of `WhereIn` is similar to go-pg, but ensure you're using the correct
	// query method calls for Bun.
	db := utils.ForDB(ctx)
	err := db.NewSelect().Model(&dbMessages).Where("id IN (?)", bun.In(keys)).Scan(ctx)
	if err != nil {
		return []*model.Message{}, []error{err}
	}
	// Mapping message IDs to messages for quick lookup.
	m := make(map[int]*model.Message)
	for _, message := range dbMessages {
		m[message.ID] = message
	}
	// Reassembling the results in the order of keys.
	messages := make([]*model.Message, len(keys))
	errs := make([]error, len(keys))
	for i, key := range keys {
		if message, ok := m[key]; ok {
			messages[i] = message
		} else {
			// Handle the case where a key does not have a corresponding user.
			errs = append(errs, fmt.Errorf("no element found for key: %d", key))
			messages[i] = nil // Keep place with nil if element is not found.
		}
	}
	return messages, errs
}
