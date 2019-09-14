package postgres

import (
	"database/sql"
	"github.com/gospeak/auth-service/model"
	"log"
)

var (
	selectToken = "SELECT id, token, user_id, created_at, updated_at, expired_at FROM tokens %s LIMIT 1"
	insertToken = "INSERT INTO tokens (token, user_id, expired_at) VALUES ($1, $2, $3) RETURNING token, user_id, expired_at"
)

type TokenStore struct {
	db *sql.DB
}

func (t *TokenStore) Get(f model.Filter) (*model.Token, error) {
	q, v := f.Filter(selectToken)
	token := &model.Token{}
	if err := t.db.
		QueryRow(q, v...).
		Scan(
			&token.ID,
			&token.Content,
			&token.UserID,
			&token.CreatedAt,
			&token.UpdatedAt,
			&token.ExpiredAt); err != nil {
		return nil, err
	}
	log.Printf("token from longe storege: %+v", token)
	return token, nil
}

// Set stet user_id via key(token) to long storage
func (t *TokenStore) Set(token *model.Token) (*model.Token, error) {
	if err := t.db.QueryRow(insertToken, token.Content, token.UserID, token.ExpiredAt).Scan(
		&token.Content,
		&token.UserID,
		&token.ExpiredAt,
	); err != nil {
		return nil, err
	}
	return token, nil
}
