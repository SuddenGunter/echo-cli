package config

import (
	"github.com/SuddenGunter/echo-cli/cmd/handler"
	"github.com/SuddenGunter/echo-cli/pkg/echo"
	"github.com/SuddenGunter/echo-cli/pkg/tokenstorage"
	"github.com/SuddenGunter/echo-cli/pkg/user"
)

// State describes app state and dependencies
type State struct {
	TokenStorage tokenstorage.TokenStorage
	Auth         handler.CobraHandler
	UserInfo     user.CurrentUserInfoProvider
	Client       echo.Client
}

func NewState(tokenStorage tokenstorage.TokenStorage, authErrorHandler func(error) error, client echo.Client) *State {
	authHandler := handler.NewAuthHandler(tokenStorage, authErrorHandler)

	return &State{
		Auth:         authHandler,
		UserInfo:     authHandler,
		TokenStorage: tokenStorage,
		Client:       client,
	}
}
