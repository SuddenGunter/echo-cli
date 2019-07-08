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
	"github.com/spf13/cobra"
)

// CobraHandler handles cobra hooks
type CobraHandler interface {
	Handle(cmd *cobra.Command, args []string) error
}

// Combine chains several command handlers and returns combined one: where all handlers will be called one-by-one
func Combine(handlers ...func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		for _, f := range handlers {
			err = f(cmd, args)
		}
		return err
	}
}
