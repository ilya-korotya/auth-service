package model

import (
	"time"
)

type User struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Email     JsonNullString `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type UserStore interface {
	Get(Filter) (*User, error)
	Set(*User) (*User, error)
}
