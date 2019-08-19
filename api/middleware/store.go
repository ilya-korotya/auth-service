package middleware

import (
	"net/http"

	"github.com/gospeak/auth-service/api/server"
	"github.com/gospeak/auth-service/postgres"
)

// CheckStore check user token in long store
func CheckStore(next server.HandlerFunc) server.HandlerFunc {
	return func(ctx server.Context, w http.ResponseWriter, r *http.Request) {
		t := r.Context().Value(Token).(string)
		tf := &postgres.TokenFilter{}
		tf.ByContent(t)
		token, err := ctx.DB.Token.Get(tf)
		// token saved in storage
		if err == nil && token.Content == t {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(ctx, w, r)
	}
}
