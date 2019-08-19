package middleware

import (
	"github.com/gospeak/auth-service/api/server"
	"net/http"
)

// Final send status code "OK" if user token is valid
func Final(ctx server.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}
