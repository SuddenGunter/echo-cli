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
	"strings"

	"github.com/SuddenGunter/echo-cli/pkg/echo"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

// createCmd represents the command for creating user in system process
var createCmd = &cobra.Command{
	Use:   "create",
	Args:  cobra.MinimumNArgs(1),
	Short: "Create user in system",
	RunE: func(cmd *cobra.Command, args []string) error {
		token, err := state.UserInfo.GetToken()
		if err != nil {
			return errors.Wrap(err, "Failed to get user token")
		}
		cl := &echo.Client{
			ConnectionString: "localhost:8080",
			Route:            "/ws",
		}
		_ = cl.Connect()
		fmt.Printf("Token: %v\n", token)
		fmt.Print("Creating user: " + strings.Join(args, ""))

		return nil
	},
}
