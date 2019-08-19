package model

// Store compositor for entities
type Store struct {
	User  UserStore
	Token TokenStore
}
