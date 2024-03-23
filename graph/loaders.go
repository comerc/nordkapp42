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
	MessageLoader  *dataloadgen.Loader[int, *model.Message]
	MemberLoader   *dataloadgen.Loader[int, *model.Member]
	ChatNameLoader *dataloadgen.Loader[int, string]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders() *Loaders {
	// define the data loaders
	return &Loaders{
		MessageLoader:  dataloadgen.NewLoader(getModels[int, model.Message], dataloadgen.WithWait(time.Millisecond)),
		MemberLoader:   dataloadgen.NewLoader(getModels[int, model.Member], dataloadgen.WithWait(time.Millisecond)),
		ChatNameLoader: dataloadgen.NewLoader(getChatNames, dataloadgen.WithWait(time.Millisecond)),
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
		m[(*dbModel).GetID()] = dbModel
	}
	// Reassembling the results in the order of keys.
	models := make([]*M, len(keys))
	errors := make([]error, len(keys))
	for i, key := range keys {
		if model, ok := m[key]; ok {
			models[i] = model
		} else {
			errors[i] = fmt.Errorf("no element found for key: %v", key)
		}
	}
	return models, errors
}

func getChatNames(ctx context.Context, keys []int) ([]string, []error) {
	type dbRow struct {
		RoomID int
		Name   string
	}
	var dbRows []*dbRow
	// Using Bun, we adjust the query method to fit Bun's API.
	// Bun's handling of `WhereIn` is similar to go-pg, but ensure you're using the correct
	// query method calls for Bun.
	db := ForDB(ctx)
	query := db.NewSelect().
		Column("room_id", "members.name").
		Table("room_members").
		Join("JOIN members ON members.id = member_id").
		Where("room_id IN (?) AND member_id != ?", bun.In(keys), ForMemberID(ctx))
	err := query.Scan(ctx, &dbRows)
	if err != nil {
		return []string{}, []error{err}
	}
	// Mapping message IDs to messages for quick lookup.
	m := make(map[int]string)
	for _, dbRow := range dbRows {
		m[dbRow.RoomID] = dbRow.Name
	}
	// Reassembling the results in the order of keys.
	names := make([]string, len(keys))
	errors := make([]error, len(keys))
	for i, key := range keys {
		if name, ok := m[key]; ok {
			names[i] = name
		} else {
			errors[i] = fmt.Errorf("no element found for key: %v", key)
		}
	}
	return names, errors
}
