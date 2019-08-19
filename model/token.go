package model

import (
	"time"
)

type Token struct {
	ID        string
	Content   string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAt time.Time
}

type TokenStore interface {
	Get(Filter) (*Token, error)
	Set(*Token) (*Token, error)
}
