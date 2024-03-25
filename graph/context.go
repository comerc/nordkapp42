package graph

import (
	"context"

	"github.com/uptrace/bun"

	"nordkapp42/pkg/jwt"
)

func GetDB(ctx context.Context) *bun.DB {
	return ctx.Value("db").(*bun.DB)
}

func GetMemberID(ctx context.Context) int {
	return jwt.GetPayload(ctx).MemberID
}
