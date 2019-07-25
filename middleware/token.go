package middleware

import (
	"context"
	"net/http"
)

type key string

const Token key = "token"

// SetToken set token to context
func SetToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get(string(Token))
		if t == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), Token, t))
		next(w, r)
	}
}
