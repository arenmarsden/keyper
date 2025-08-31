package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newBackupCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "backup",
		Short: "Backup a file or directly immediately, encryption is automatically applied",
		RunE: func(_ *cobra.Command, _ []string) error {
			fmt.Println("hi")
			return nil
		},
	}
}
