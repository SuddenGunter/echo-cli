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
package cmd

import (
	"log"

	"github.com/SuddenGunter/echo-cli/cmd/handler"

	"github.com/SuddenGunter/echo-cli/cmd/config"

	"github.com/spf13/cobra"
)

var (
	tokenFlag string
	state     *config.State
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "echo-cli",
	Short: "Echo-CLI app created as one-evening experiment where main goal was get some skills with Cobra",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(initialState *config.State) {
	state = initialState

	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(createCmd)
	createCmd.RunE = handler.Combine(state.Auth.Handle, createCmd.RunE)

	rootCmd.AddCommand(authCmd)
	authCmd.Flags().StringVarP(&tokenFlag, "token", "t", "", "Base64 encoded auth token value")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
