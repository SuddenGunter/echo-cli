/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"fmt"

	"github.com/SuddenGunter/echo-cli/pkg/tokenstorage"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var ErrEmptyAuthToken = errors.New("Auth token must be not empty")

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:     "auth",
	Short:   "Authorize local user to echo-server",
	Example: "echo-cli auth -t=SECURITY_TOKEN",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(tokenFlag) <= 0 {
			return ErrEmptyAuthToken
		}
		err := state.TokenStorage.Save(tokenFlag)
		if err != nil {
			return errors.Wrap(err, "Failed save auth token on auth command")
		}
		return nil
	},
}

func UnauthorizedErrorHandler(err error) error {
	if err == tokenstorage.ErrTokenNotFound {
		fmt.Println("User must be authorized before using this command")
		_ = authCmd.Help()
		return nil
	}
	return errors.Wrap(err, "Failed to read auth token on user creation")
}
