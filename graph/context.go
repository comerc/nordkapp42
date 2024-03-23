package graph

import (
	"context"

	"github.com/uptrace/bun"
)

func ForDB(ctx context.Context) *bun.DB {
	return ctx.Value("db").(*bun.DB)
}

func ForMemberID(ctx context.Context) int {
	return ctx.Value("memberID").(int)
}
