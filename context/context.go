package context

import (
	"database/sql"

	"github.com/go-redis/redis"
)

// Context contains data for work middleware and handlers
type Context struct {
	SessinID string
	Cache    *redis.Client
	Store    *sql.DB
}
