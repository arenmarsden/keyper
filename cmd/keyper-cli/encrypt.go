package main

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

func newEncryptCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "encrypt",
		Short: "Manually encrypt files using your stored key",
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("please specify a file name(s)")
			}
			for _, fileName := range args {
				log.Printf("encrypting %s\n", fileName)
			}
			return nil
		},
	}
}
