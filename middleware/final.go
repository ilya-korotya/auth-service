package middleware

import (
	"net/http"
)

// Final send status code "OK" if user token is valid
func Final(ctx Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}
