package middleware

import (
	"net/http"

	"github.com/gospeak/auth-service/context"
	"github.com/gospeak/auth-service/server"
)

// SetToken set token to context
func SetToken(next server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("token")
		if t == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx.SessinID = t
		next(ctx, w, r)
	}
}
