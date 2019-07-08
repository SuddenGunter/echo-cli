package handler

import (
	"github.com/SuddenGunter/echo-cli/pkg/tokenstorage"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type Auth struct {
	storage      tokenstorage.TokenStorage
	token        string
	errorHandler func(error) error
}

// GetToken returns user's token if it exists
func (handler *Auth) GetToken() (string, error) {
	if len(handler.token) <= 0 {
		return "", errors.New("Error token must not be empty")
	}
	return handler.token, nil
}

// NewAuthHandler creates new auth handler instance
// errorHandler triggers if authorization failed (e.g. token not found)
func NewAuthHandler(tokenStorage tokenstorage.TokenStorage, errorHandler func(error) error) *Auth {
	return &Auth{storage: tokenStorage,
		errorHandler: errorHandler,
	}
}

// Handle handles user's authorization and can be used in handlers pipeline
func (handler *Auth) Handle(cmd *cobra.Command, args []string) error {
	token, err := handler.storage.Read()
	if err != nil {
		return handler.errorHandler(err)
	}
	handler.token = token
	return nil
}
