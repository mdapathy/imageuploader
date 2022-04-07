package httpapi

import "context"

type ContextKey string

const (
	CtxKeyUserID ContextKey = "userID"
)

func UserIDToContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, CtxKeyUserID, id)
}

// UserIDFromContext returns user id from the context.
func UserIDFromContext(ctx context.Context) string {
	id, ok := ctx.Value(CtxKeyUserID).(string)
	if !ok {
		return ""
	}

	return id
}
