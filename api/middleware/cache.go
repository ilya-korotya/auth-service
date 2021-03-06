package middleware

import (
	"github.com/gospeak/auth-service/api/server"
	"net/http"
)

// CacheCheck check user token in cache
func CacheCheck(next server.HandlerFunc) server.HandlerFunc {
	return func(ctx server.Context, w http.ResponseWriter, r *http.Request) {
		t := r.Context().Value(Token).(string)
		u := ctx.Cache.Get(t)
		// if user is exist in cache - authorize this user
		if u != "" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// if user don't in cache - check him in long storage
		next(ctx, w, r)
	}
}
