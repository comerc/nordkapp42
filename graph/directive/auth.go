package directive

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"

	"nordkapp42/graph"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	memberID := graph.ForMemberID(ctx)
	if memberID != 0 {
		return next(ctx)
	} else {
		return nil, errors.New("Unauthorised")
	}
}
