package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"nordkapp42/graph/model"
)

// Member is the resolver for the member field.
func (r *messageResolver) Member(ctx context.Context, obj *model.Message) (*model.Member, error) {
	loader := ForLoaders(ctx).MemberLoader
	return loader.Load(ctx, obj.MemberID)
}

// Rooms is the resolver for the rooms field.
func (r *queryResolver) Rooms(ctx context.Context) ([]*model.Room, error) {
	db := ForDB(ctx)
	var rooms []*model.Room
	query := db.NewSelect().
		Column("rooms.*").
		Table("rooms").
		Join("JOIN room_members AS t").
		JoinOn("t.room_id = rooms.id").
		JoinOn("t.member_id = ?", ForMemberID(ctx))
	// if limit != nil {
	// 	query = query.Limit(*limit)
	// }
	// if offset != nil {
	// 	query = query.Offset(*offset)
	// }
	if err := query.Scan(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

// Name is the resolver for the name field.
func (r *roomResolver) Name(ctx context.Context, room *model.Room) (string, error) {
	if room.Kind == model.RoomKindEnumChat {
		loader := ForLoaders(ctx).ChatNameLoader
		return loader.Load(ctx, room.ID)
	}
	return room.Name, nil
}

// Messages is the resolver for the messages field.
func (r *roomResolver) Messages(ctx context.Context, obj *model.Room) ([]*model.Message, error) {
	messageIDs := []int{2, 3}
	loader := ForLoaders(ctx).MessageLoader
	messages, errorSlice := loader.LoadAll(ctx, messageIDs)
	return messages, errorSlice
}

// Rooms is the resolver for the rooms field.
func (r *subscriptionResolver) Rooms(ctx context.Context) (<-chan []*model.Room, error) {
	panic(fmt.Errorf("not implemented: Rooms - rooms"))
}

// Message returns MessageResolver implementation.
func (r *Resolver) Message() MessageResolver { return &messageResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Room returns RoomResolver implementation.
func (r *Resolver) Room() RoomResolver { return &roomResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type messageResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type roomResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
