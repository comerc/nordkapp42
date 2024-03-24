package graph

import (
	"context"

	"github.com/uptrace/bun"
)

func ForDB(ctx context.Context) *bun.DB {
	return ctx.Value("db").(*bun.DB)
}

func ForMemberID(ctx context.Context) int {
	res, dummy := ctx.Value("memberID").(int)
	_ = dummy
	// без dummy при отсутствии memberID:
	// "interface conversion: interface {} is nil, not int"
	return res
}
