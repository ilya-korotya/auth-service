package postgres

import (
	"fmt"
	"log"
	"strings"
)

type UserFilter struct {
	ID    string
	Name  string
	Email string
	Token string
}

// ByID add to query WHERE part by user ID
func (u *UserFilter) ByID(id string) *UserFilter {
	u.ID = id
	return u
}

func (u *UserFilter) ByName(name string) *UserFilter {
	u.Name = name
	return u
}

func (u *UserFilter) ByEmail(email string) *UserFilter {
	u.Email = email
	return u
}

func (u *UserFilter) ByToken(token string) *UserFilter {
	u.Token = token
	return u
}

func (u *UserFilter) Filter(sql string) (string, []interface{}) {
	var values []interface{}
	var columns []string
	var query string
	if u.ID != "" {
		columns = append(columns, "id")
		values = append(values, u.ID)
	}
	if u.Name != "" {
		columns = append(columns, "user_name")
		values = append(values, u.Name)
	}
	if u.Email != "" {
		columns = append(columns, "email")
		values = append(values, u.Email)
	}
	if u.Token != "" {
		columns = append(columns, "tokens.token")
		values = append(values, u.Token)
		query += "INNER JOIN tokens ON (tokens.user_id = users.id) "
	}
	if len(values) > 0 {
		query += "WHERE "
	}
	for i := range values {
		query += fmt.Sprintf("%s=$%v AND ", columns[i], i+1)
	}
	if len(values) > 0 {
		query = strings.TrimSuffix(query, " AND ")
	}
	log.Println("user query:", query)
	return fmt.Sprintf(sql, query), values
}
