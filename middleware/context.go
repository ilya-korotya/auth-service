package middleware

import (
	"net/http"
	"time"
)

// GetToken return token as string from store
type GetToken interface {
	Get(string) string
}

// SetToken to store with expiration date
type SetToken interface {
	Set(string, string, time.Duration) error
}

// Cache for cache(redis) provider
type Cache interface {
	GetToken
	SetToken
}

// LongStorage for database(posthres) provider
type LongStorage interface {
	GetToken
	SetToken
}

// Context for app. Contains logic work for user token
type Context struct {
	Cache Cache
	Store LongStorage
}

// HandlerFunc discare handler type
type HandlerFunc func(ctx Context, w http.ResponseWriter, r *http.Request)
