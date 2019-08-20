package main

import (
	"github.com/gospeak/auth-service/api/inner"
	"github.com/gospeak/auth-service/api/public"
	"github.com/gospeak/auth-service/api/server"
	"github.com/gospeak/auth-service/model"
	"github.com/gospeak/auth-service/postgres"
	"log"

	"github.com/go-redis/redis"
	"github.com/gospeak/auth-service/api/middleware"
	"github.com/gospeak/auth-service/cache"
	_ "github.com/lib/pq"
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
	s, err := postgres.New("postgres://auth_user_role:@auth-postgres:5432/auth?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to store:", err)
	}
	ctx := server.Context{
		Cache: c,
		DB: &model.Store{
			User:  s.UserStore,
			Token: s.TokenStore,
		},
	}

	// run public api
	m := middleware.InitToken(middleware.CacheCheck(middleware.CheckStore(middleware.Final)))
	pub := server.New(ctx)
	pub.Get("/", m) // must be in all urls
	pub.Post("/user/registration", public.RegistrationUser)
	go func() {
		log.Println("run public api")
		log.Println("public apu error:", pub.Run(":8080"))
	}()

	// run inner api
	in := server.New(ctx)
	in.Get("/v1/user", middleware.InitToken(inner.GetUser))
	log.Println("run internal api")
	log.Println("inner api error:", in.Run(":8081"))
}
