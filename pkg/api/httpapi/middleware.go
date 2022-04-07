package httpapi

import (
	"errors"
	"net/http"
)

const (
	HeaderUserID string = "X-User-Id"
)

func UserMiddleware(next http.Handler) http.Handler {
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		userID := r.Header.Get(HeaderUserID)
		if userID == "" {
			return errors.New("expected user id")
		}
		ctx = UserIDToContext(ctx, userID)

		next.ServeHTTP(w, r.WithContext(ctx))

		return nil
	})
}
