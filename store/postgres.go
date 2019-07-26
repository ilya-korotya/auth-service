package store

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// Config path to database connect
type Config string

// Client contains connect to database
type Client struct {
	db *sql.DB
}

// Open new connect to database
func Open(c Config) (*Client, error) {
	db, err := sql.Open("postgres", string(c))
	return &Client{
		db: db,
	}, err
}

// Get return user_id form long storage
func (c Client) Get(key string) string {
	var res string
	// TODO (ilya-korotya): check experation time in SQL request?
	err := c.db.QueryRow("SELECT user_id from tokens WHERE id=$1", key).Scan(&res)
	if err != nil {
		return ""
	}
	return res
}

// Set stet user_id via key(token) to long storage
func (c Client) Set(key, value string, expiration time.Duration) error {
	return nil
}
