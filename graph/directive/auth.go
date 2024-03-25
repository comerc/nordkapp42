package directive

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"

	"nordkapp42/pkg/jwt"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	payload := jwt.GetPayload(ctx)
	if jwt.IsExpired(payload) {
		return nil, errors.New("JWT was expired")
	}
	if payload.MemberID == 0 {
		return nil, errors.New("Unauthorised")
	}
	return next(ctx)
}
