package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	// memberID := ctx.Value("memberID")
	// fmt.Println("*************")
	// fmt.Printf("%#v\n", memberID)
	// fmt.Println("*************")
	// // if memberID != nil {
	// // return next(ctx)
	// // } else {
	// return nil, errors.New("Unauthorised")
	// }
	return next(ctx)
}
