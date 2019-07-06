package handler

import (
	"github.com/spf13/cobra"
)

// CobraHandler handles cobra hooks
type CobraHandler interface {
	Handle(cmd *cobra.Command, args []string) error
}

func Combine(handlers ...func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		for _, f := range handlers {
			err = f(cmd, args)
		}
		return err
	}
}
