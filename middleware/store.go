package middleware

import (
	"net/http"
)

// CheckStore check user token in long store
func CheckStore(next HandlerFunc) HandlerFunc {
	return func(ctx Context, w http.ResponseWriter, r *http.Request) {
		t := r.Context().Value(Token).(string)
		u := ctx.Store.Get(t)
		// if user is exist in cache - authorize this user
		if u != "" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// if user don't in cache - check him in long storage
		next(ctx, w, r)
	}
}
