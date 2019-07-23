package middleware

import (
	"net/http"

	"github.com/gospeak/auth-service/context"
)

// Final send status code "OK" if user token is valid
func Final(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
