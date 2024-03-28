package graph

// TODO: нужен канал синхронизации (очистки) кеша лоадеров через нотификацию изменений моделей

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/vikstrous/dataloadgen"

	"nordkapp42/graph/model"
)

func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value("loaders").(*Loaders)
}

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	ManyMessagesLoader    *dataloadgen.Loader[int, []*model.Message]
	MemberLoader          *dataloadgen.Loader[int, *model.Member]
	RoomLoader            *dataloadgen.Loader[int, *model.Room]
	ChatRoomPropsLoader   *dataloadgen.Loader[int, *model.RoomProps]
	CommonRoomPropsLoader *dataloadgen.Loader[int, *model.RoomProps]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders() *Loaders {
	// define the data loaders
	return &Loaders{
		ManyMessagesLoader: dataloadgen.NewLoader(
			fetchManyMessages,
			dataloadgen.WithWait(time.Millisecond),
		),
		MemberLoader: dataloadgen.NewLoader(
			fetchModels(getDBModels[int, model.Member]("id")),
			dataloadgen.WithWait(time.Millisecond),
		),
		RoomLoader: dataloadgen.NewLoader(
			fetchModels(getDBModels[int, model.Room]("id")),
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

func fetchManyMessages(ctx context.Context, keys []int) ([][]*model.Message, []error) {
	dbModels, err := getDBModels[int, model.Message]("room_id")(ctx, keys)
	if err != nil {
		return [][]*model.Message{}, []error{err}
	}
	m := make(map[int][]*model.Message)
	for _, dbModel := range dbModels {
		m[dbModel.RoomID] = append(m[dbModel.RoomID], dbModel)
	}
	models := make([][]*model.Message, len(keys))
	for i, key := range keys {
		if model, ok := m[key]; ok {
			models[i] = model
		}
	}
	return models, nil
}

type FetchModelsFn[T comparable, M model.Model[T]] func(ctx context.Context, keys []T) ([]*M, []error)
type GetDBModelsFn[T comparable, M model.Model[T]] func(ctx context.Context, keys []T) ([]*M, error)

func fetchModels[T comparable, M model.Model[T]](getDBModels GetDBModelsFn[T, M]) FetchModelsFn[T, M] {
	return func(ctx context.Context, keys []T) ([]*M, []error) {
		dbModels, err := getDBModels(ctx, keys)
		if err != nil {
			return dbModels, []error{err}
		}
		// Mapping model IDs to models for quick lookup.
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
		db := GetDB(ctx)
		query := db.NewSelect().
			Model(&dbModels).
			Where("? IN (?)", bun.Ident(idName), bun.In(keys))
		err := query.Scan(ctx)
		return dbModels, err
	}
}

func getChatRoomProps[T comparable, M model.Model[T]](ctx context.Context, keys []T) ([]*M, error) {
	var dbModels []*M
	db := GetDB(ctx)
	query := db.NewSelect().
		Column("room_id", "members.name").
		Table("room_members").
		Join("JOIN members ON members.id = member_id").
		Where("room_id IN (?) AND member_id != ?", bun.In(keys), GetMemberID(ctx))
	err := query.Scan(ctx, &dbModels)
	return dbModels, err
}
