package middleware

import (
	"context"
	"github.com/gospeak/auth-service/api/server"
	"net/http"
)

type key string

// Token name for request token
const Token key = "token"

// InitToken init token to requst context
func InitToken(next server.HandlerFunc) server.HandlerFunc {
	return func(ctx server.Context, w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get(string(Token))
		if t == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), Token, t))
		next(ctx, w, r)
	}
}
