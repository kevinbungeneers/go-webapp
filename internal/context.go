package internal

import "context"

type contextKey int

const (
	requestIdKey contextKey = iota
)

func NewRequestIdContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIdKey, id)
}

func RequestIdFromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(requestIdKey).(string)
	return v, ok
}
