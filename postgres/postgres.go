package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Config path to database connect
type Config string

// Store contains connect to database
type Store struct {
	UserStore  *UserStore
	TokenStore *TokenStore
}

// New new connect to database
func New(c Config) (*Store, error) {
	db, err := sql.Open("postgres", string(c))
	return &Store{
		UserStore:  &UserStore{db: db},
		TokenStore: &TokenStore{db: db},
	}, err
}
