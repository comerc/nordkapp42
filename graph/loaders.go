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
	MessageLoader         *dataloadgen.Loader[int, *model.Message]
	MemberLoader          *dataloadgen.Loader[int, *model.Member]
	ChatRoomPropsLoader   *dataloadgen.Loader[int, *model.RoomProps]
	CommonRoomPropsLoader *dataloadgen.Loader[int, *model.RoomProps]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders() *Loaders {
	// define the data loaders
	return &Loaders{
		MessageLoader: dataloadgen.NewLoader(
			fetchModels(getDBModels[int, model.Message]("id")),
			dataloadgen.WithWait(time.Millisecond),
		),
		MemberLoader: dataloadgen.NewLoader(
			fetchModels(getDBModels[int, model.Member]("id")),
			dataloadgen.WithWait(time.Millisecond),
		),
		CommonRoomPropsLoader: dataloadgen.NewLoader(
			fetchModels(getDBModels[int, model.RoomProps]("room_id")),
			dataloadgen.WithWait(time.Millisecond),
		),
		ChatRoomPropsLoader: dataloadgen.NewLoader(
			fetchModels(getChatRoomProps[int, model.RoomProps]),
			dataloadgen.WithWait(time.Millisecond),
		),
	}
}

type FetchModelsFn[T comparable, M model.Model[T]] func(ctx context.Context, keys []T) ([]*M, []error)
type GetDBModelsFn[T comparable, M model.Model[T]] func(ctx context.Context, keys []T) ([]*M, error)

func fetchModels[T comparable, M model.Model[T]](getDBModels GetDBModelsFn[T, M]) FetchModelsFn[T, M] {
	return func(ctx context.Context, keys []T) ([]*M, []error) {
		dbModels, err := getDBModels(ctx, keys)
		if err != nil {
			return dbModels, []error{err}
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
}

func getDBModels[T comparable, M model.Model[T]](idName string) GetDBModelsFn[T, M] {
	return func(ctx context.Context, keys []T) ([]*M, error) {
		var dbModels []*M
		db := ForDB(ctx)
		// Using Bun, we adjust the query method to fit Bun's API.
		// Bun's handling of `WhereIn` is similar to go-pg, but ensure you're using the correct
		// query method calls for Bun.
		query := db.NewSelect().Model(&dbModels).Where("? IN (?)", bun.Ident(idName), bun.In(keys))
		err := query.Scan(ctx)
		return dbModels, err
	}
}

func getChatRoomProps[T comparable, M model.Model[T]](ctx context.Context, keys []T) ([]*M, error) {
	var dbModels []*M
	db := ForDB(ctx)
	// Using Bun, we adjust the query method to fit Bun's API.
	// Bun's handling of `WhereIn` is similar to go-pg, but ensure you're using the correct
	// query method calls for Bun.
	query := db.NewSelect().
		Column("room_id", "members.name").
		Table("room_members").
		Join("JOIN members ON members.id = member_id").
		Where("room_id IN (?) AND member_id != ?", bun.In(keys), ForMemberID(ctx))
	err := query.Scan(ctx, &dbModels)
	return dbModels, err
}
