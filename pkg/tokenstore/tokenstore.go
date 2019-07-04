package tokenstore

import "github.com/pkg/errors"

// TokenStore defines interface for auth token persistent storage
type TokenStore interface {

	// Save takes token string and save it somewhere
	Save(token string) error
	// Read and return token from store
	// Returns ErrTokenNotFound if token not found
	Read() (string, error)
}

var ErrTokenNotFound = errors.New("Auth token not found")
