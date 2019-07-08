/*
Copyright © 2019 ARTEM KOLOMYTSEV kolomytsev1996@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
