package main

import "github.com/spf13/cobra"

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "keyper-cli",
		Long:  "keyper-cli allows for manual intervention of the keyper deameon, manual backup and encryption",
		Short: "keyper-cli allows you to manually backup and encrypt your files at will",
	}
}
