package mock

import "time"

// GetSetMock mock for cache provider
type GetSetMock struct {
	GetIsCall bool
	GetFunc   func(string) string
	SetIsCall bool
	SetFunc   func(string, string, time.Duration) error
}

func (c *GetSetMock) Get(key string) string {
	c.GetIsCall = true
	return c.GetFunc(key)
}

func (c *GetSetMock) Set(key, value string, expiration time.Duration) error {
	c.SetIsCall = true
	return c.SetFunc(key, value, expiration)
}
