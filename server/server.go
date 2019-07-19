package server

import (
	"net/http"

	"github.com/gospeak/auth-service/context"
)

type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func ListenAndServe(ctx context.Context, addr string, h HandlerFunc) error {
	return http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h(ctx, w, r)
	}))
}
