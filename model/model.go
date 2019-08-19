package model

// Model compositor for entities
type Store struct {
	User  UserStore
	Token TokenStore
}
