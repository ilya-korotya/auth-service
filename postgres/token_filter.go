package postgres

import (
	"fmt"
	"log"
	"strings"
)

type TokenFilter struct {
	ID      string
	Content string
	UserID  string
}

func (t *TokenFilter) ByID(id string) *TokenFilter {
	t.ID = id
	return t
}

func (t *TokenFilter) ByContent(content string) *TokenFilter {
	t.Content = content
	return t
}

func (t *TokenFilter) ByUserID(userID string) *TokenFilter {
	t.UserID = userID
	return t
}

func (t *TokenFilter) Filter(sql string) (string, []interface{}) {
	var values []interface{}
	var columns []string
	var query string
	if t.ID != "" {
		columns = append(columns, "id")
		values = append(values, t.ID)
	}
	if t.Content != "" {
		columns = append(columns, "token")
		values = append(values, t.Content)
	}
	if t.UserID != "" {
		columns = append(columns, "user_id")
		values = append(values, t.UserID)
		query += "INNER JOIN users ON (tokens.user_id = users.id) "
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
	log.Println("token query:", query)
	return fmt.Sprintf(sql, query), values
}
