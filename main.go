package main

import (
	"log"

	"github.com/gospeak/auth-service/context"
	"github.com/gospeak/auth-service/middleware"
	"github.com/gospeak/auth-service/server"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Context{}
	h := middleware.SetToken(middleware.Final)
	log.Fatal("cannot start listener:", server.ListenAndServe(ctx, ":8080", h))
}
