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
		MessageLoader:         dataloadgen.NewLoader(getModels[int, model.Message]("id"), dataloadgen.WithWait(time.Millisecond)),
		MemberLoader:          dataloadgen.NewLoader(getModels[int, model.Member]("id"), dataloadgen.WithWait(time.Millisecond)),
		ChatRoomPropsLoader:   dataloadgen.NewLoader(getChatRoomProps, dataloadgen.WithWait(time.Millisecond)),
		CommonRoomPropsLoader: dataloadgen.NewLoader(getModels[int, model.RoomProps]("room_id"), dataloadgen.WithWait(time.Millisecond)),
	}
}

func getModels[T comparable, M model.Model[T]](idIdent string) func(ctx context.Context, keys []T) ([]*M, []error) {
	return func(ctx context.Context, keys []T) ([]*M, []error) {
		var dbModels []*M
		// Using Bun, we adjust the query method to fit Bun's API.
		// Bun's handling of `WhereIn` is similar to go-pg, but ensure you're using the correct
		// query method calls for Bun.
		db := ForDB(ctx)
		err := db.NewSelect().Model(&dbModels).Where("? IN (?)", bun.Ident(idIdent), bun.In(keys)).Scan(ctx)
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
}

func getChatRoomProps(ctx context.Context, keys []int) ([]*model.RoomProps, []error) {
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
		return []*model.RoomProps{}, []error{err}
	}
	// Mapping message IDs to messages for quick lookup.
	m := make(map[int]*dbRow)
	for _, dbRow := range dbRows {
		m[dbRow.RoomID] = dbRow
	}
	// Reassembling the results in the order of keys.
	props := make([]*model.RoomProps, len(keys))
	errors := make([]error, len(keys))
	for i, key := range keys {
		if dbRow, ok := m[key]; ok {
			props[i] = &model.RoomProps{Name: dbRow.Name}
		} else {
			errors[i] = fmt.Errorf("no element found for key: %v", key)
		}
	}
	return props, errors
}

// func getCommonRoomProps(ctx context.Context, keys []int) ([]*model.RoomProps, []error) {
// 	type dbRow struct {
// 		RoomID int
// 		Name   string
// 	}
// 	var dbRows []*dbRow
// 	// Using Bun, we adjust the query method to fit Bun's API.
// 	// Bun's handling of `WhereIn` is similar to go-pg, but ensure you're using the correct
// 	// query method calls for Bun.
// 	db := ForDB(ctx)
// 	query := db.NewSelect().
// 		Column("room_id", "members.name").
// 		Table("room_members").
// 		Join("JOIN members ON members.id = member_id").
// 		Where("room_id IN (?) AND member_id != ?", bun.In(keys), ForMemberID(ctx))
// 	err := query.Scan(ctx, &dbRows)
// 	if err != nil {
// 		return []*model.RoomProps{}, []error{err}
// 	}
// 	// Mapping message IDs to messages for quick lookup.
// 	m := make(map[int]*dbRow)
// 	for _, dbRow := range dbRows {
// 		m[dbRow.RoomID] = dbRow
// 	}
// 	// Reassembling the results in the order of keys.
// 	props := make([]*model.RoomProps, len(keys))
// 	errors := make([]error, len(keys))
// 	for i, key := range keys {
// 		if dbRow, ok := m[key]; ok {
// 			props[i] = &model.RoomProps{Name: dbRow.Name}
// 		} else {
// 			errors[i] = fmt.Errorf("no element found for key: %v", key)
// 		}
// 	}
// 	return props, errors
// }
