package postgres

import (
	"database/sql"

	"github.com/gospeak/auth-service/model"
)

var (
	selectUser  = "SELECT users.id, user_name, email, password, users.created_at, users.updated_at FROM users %s LIMIT 1"
	insertUser  = "INSERT INTO users (user_name, email, password) VALUES ($1, $2, $3) RETURNING id, user_name, email, password"
	insertToken = "INSERT INTO tokens (token, user_id, expired_at) VALUES ($1, $2, $3) RETURNING id, token, user_id, expired_at"
)

type UserStore struct {
	db *sql.DB
}

func (u *UserStore) Get(f model.Filter) (*model.User, error) {
	q, v := f.Filter(selectUser)
	user := &model.User{}
	if err := u.db.
		QueryRow(q, v...).
		Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserStore) Set(user *model.User) (*model.User, error) {
	tx, err := u.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// rewrite user after registration
	if err := tx.QueryRow(insertUser, user.Name, user.Email, user.Password).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.QueryRow(insertToken, user.Token.Content, user.ID, user.Token.ExpiredAt).Scan(
		&user.Token.ID,
		&user.Token.Content,
		&user.Token.UserID,
		&user.Token.ExpiredAt,
	); err != nil {
		tx.Rollback()
		return nil, err
	}
	return user, tx.Commit()
}
