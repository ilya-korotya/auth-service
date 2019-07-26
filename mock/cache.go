package mock

import "time"

// CacheMock mock for cache provider
type CacheMock struct {
	GetIsCall bool
	GetFunc   func(string) string
	SetIsCall bool
	SetFunc   func(string, string, time.Duration) error
}

func (c *CacheMock) Get(key string) string {
	c.GetIsCall = true
	return c.GetFunc(key)
}

func (c *CacheMock) Set(key, value string, expiration time.Duration) error {
	c.SetIsCall = true
	return c.SetFunc(key, value, expiration)
}
