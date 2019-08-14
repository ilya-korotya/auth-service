package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gospeak/auth-service/cache"
	"github.com/gospeak/auth-service/middleware"
	"github.com/gospeak/auth-service/store"
)

func main() {
	// TODO (ilya-korotya): maybe better use URL for connection
	c, err := cache.Open(&redis.Options{
		Addr:     "auth-redis:6379",
		Password: "",
		DB:       0,
	})
	if err != nil {
		log.Fatal("cannot connect to cache:", err)
	}
	s, err := store.Open("postgres://auth_user_role:@auth-postgres:5432/auth?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to store:", err)
	}
	ctx := middleware.Context{
		Cache: c,
		Store: s,
	}
	m := middleware.InitToken(middleware.CacheCheck(middleware.CheckStore(middleware.Final)))
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m(ctx, w, r)
	}))
}
