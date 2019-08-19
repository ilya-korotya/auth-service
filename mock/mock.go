package mock

import (
	"github.com/gospeak/auth-service/model"
	"time"
)

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

// UserStoreMock mock for UserStore
type UserStoreMock struct {
	GetIsCall bool
	GetFunc   func(model.Filter) (*model.User, error)
	SetIsCall bool
	SetFunc   func(*model.User) (*model.User, error)
}

func (u *UserStoreMock) Get(f model.Filter) (*model.User, error) {
	u.GetIsCall = true
	return u.GetFunc(f)
}

func (u *UserStoreMock) Set(user *model.User) (*model.User, error) {
	u.SetIsCall = true
	return u.SetFunc(user)
}

// TokenStoreMock mock for TokenStore
type TokenStoreMock struct {
	GetIsCall bool
	GetFunc   func(model.Filter) (*model.Token, error)
	SetIsCall bool
	SetFunc   func(*model.Token) (*model.Token, error)
}

func (t *TokenStoreMock) Get(f model.Filter) (*model.Token, error) {
	t.GetIsCall = true
	return t.GetFunc(f)
}

func (t *TokenStoreMock) Set(token *model.Token) (*model.Token, error) {
	t.SetIsCall = true
	return t.SetFunc(token)
}
