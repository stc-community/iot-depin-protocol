package types

import (
	"context"
)

type creatorCtx struct{}

// NewCreator 创建者的上下文
func NewCreator(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, creatorCtx{}, userID)
}

// FromCreator 从上下文中获取创建者
func FromCreator(ctx context.Context) (string, bool) {
	v := ctx.Value(creatorCtx{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s, s != ""
		}
	}
	return "", false
}
