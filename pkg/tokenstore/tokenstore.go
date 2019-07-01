package tokenstore

import "errors"

// TokenStore defines interface for auth token persistent storage
type TokenStore interface {

	// Save takes token string and save in in some persistent storage
	Save(token string) error
	// Read token from storage
	Read() (string, error)
}

var ErrTokenNotFound = errors.New("Token not found")