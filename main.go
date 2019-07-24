package main

import (
	"net/http"

	"github.com/gospeak/auth-service/middleware"
	_ "github.com/lib/pq"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(middleware.SetToken(middleware.Final)))
}
