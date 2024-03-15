package utils

import (
	"context"

	"github.com/uptrace/bun"
)

func ForDB(ctx context.Context) *bun.DB {
	return ctx.Value("db").(*bun.DB)
}
