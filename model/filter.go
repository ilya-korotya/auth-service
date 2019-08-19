package model

// Filter generate sql and data for request
type Filter interface {
	Filter(string) (string, []interface{})
}
