package tokenstorage

import "github.com/pkg/errors"

// TokenStorage defines interface for auth token persistent storage
type TokenStorage interface {

	// Delete old tokens and save new
	Save(token string) error
	// Read and return token from storage
	// Returns ErrTokenNotFound if token not found
	Read() (string, error)
}

var ErrTokenNotFound = errors.New("Auth token not found")
