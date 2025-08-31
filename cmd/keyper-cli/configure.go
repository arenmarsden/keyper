package main

import (
	"log"

	"github.com/spf13/cobra"
)

func newConfigureCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "configure",
		Short: "Configure keyper to run headless and specify an encryption key",
		RunE: func(_ *cobra.Command, _ []string) error {
			// todo: make this a prompt
			log.Println("this will be a prompt")
			return nil
		},
	}
}
