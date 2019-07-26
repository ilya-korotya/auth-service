package middleware

import (
	"context"
	"net/http"
)

type key string

// Token name for request token
const Token key = "token"

// InitToken init token to requst context
func InitToken(next HandlerFunc) HandlerFunc {
	return func(ctx Context, w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get(string(Token))
		if t == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), Token, t))
		next(ctx, w, r)
	}
}
