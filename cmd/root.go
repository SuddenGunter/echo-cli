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
  "github.com/SuddenGunter/echo-cli/pkg/tokenstore"
  "github.com/spf13/cobra"
  "os"
)

var (
  tokenStore tokenstore.TokenStore
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "echo-cli",
  Short: "Echo-CLI app created as one-evening experiment where main goal was get some skills with Cobra",
  PreRunE: configureDependencies,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

// configureDependencies configures all required dependencies
func configureDependencies(cmd *cobra.Command, args []string) error {
  tokenStore = tokenstore.NewTempFileTokenStore(tokenstore.DefaultTempFileTokenStoreConfig)
  //todo check if file exists
  return nil
}


