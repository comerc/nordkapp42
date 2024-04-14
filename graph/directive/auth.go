package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	// payload := jwt.GetPayload(ctx)
	// if payload.IsExpired() {
	// 	return nil, errors.New("JWT was expired")
	// }
	// if payload.MemberID == 0 {
	// 	return nil, errors.New("Unauthorised")
	// }
	return next(ctx)
}

// type ckey string
// func User(ctx context.Context, obj interface{}, next graphql.Resolver, username string) (res interface{}, err error) {
// 	return next(context.WithValue(ctx, ckey("username"), username))
// }
