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
	"os"

	"github.com/spf13/cobra"
)

// userCmd represents the user command, which is only a glue
// for all user-related commands like 'create'
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SetOut(os.Stderr)
		cmd.HelpFunc()(cmd, args)
		return nil
	},
}
