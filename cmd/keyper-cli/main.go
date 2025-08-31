package main

import "log"

func main() {
	rootCmd := newRootCmd()
	rootCmd.AddCommand(
		newBackupCmd(),
		newEncryptCmd(),
		newConfigureCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("an error occurred while running the command: %+v", err)
	}
}
