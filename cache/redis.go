package cache

import (
	"time"

	"github.com/go-redis/redis"
)

// Client use access to cache
type Client struct {
	r *redis.Client
}

// Open connection to cache provider
func Open(c *redis.Options) (*Client, error) {
	rc := redis.NewClient(c)
	return &Client{rc}, rc.Ping().Err()
}

// Get string by key from cache provider
func (c Client) Get(key string) string {
	res, err := c.r.Get(key).Result()
	if err != nil {
		return ""
	}
	return res
}

// Set string to cache. If expiration 0 cache don't remove by time
func (c Client) Set(key, value string, expiration time.Duration) error {
	return c.r.Set(key, value, expiration).Err()
}
