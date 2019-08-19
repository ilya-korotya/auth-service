package postgres

import (
	"database/sql"

	"github.com/gospeak/auth-service/model"
)

var (
	selectUser = "SELECT users.id, user_name, email, password, users.created_at, users.updated_at FROM users %s LIMIT 1"
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
	return nil, nil
}
