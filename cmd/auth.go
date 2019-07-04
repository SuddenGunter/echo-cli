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
	"errors"
	"github.com/spf13/cobra"
)

var ErrEmptyAuthToken = errors.New("Auth token must be not empty")

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:     "auth",
	Short:   "Authorize local user to echo-server",
	Example: "echo-cli auth -t=SECURITY_TOKEN",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(token) <= 0 {
			return ErrEmptyAuthToken
		}
		err := tokenStorage.Save(token)
		if err != nil {
			return err
		}
		return nil
	},
}
