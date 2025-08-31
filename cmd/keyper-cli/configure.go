package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/arenmarsden/keyper/internal/config"
)

func newConfigureCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "configure",
		Short: "Configure keyper to run headless and specify an encryption key",
		RunE: func(_ *cobra.Command, _ []string) error {
			cfg, err := config.LoadConfig()
			if err != nil {
				return err
			}

			cfg.StorageProvider = promptInput("Storage Provider")
			cfg.Endpoint = promptInput("Endpoint (e.g https://s3.amazonaws.com)")
			cfg.AccessKeyID = promptInput("Access Key Id")
			cfg.Region = promptInputOptional("Region (optional)")
			cfg.UseSSL = promptBool("Use SSL? (true, false)", false)

			if err := config.WriteConfig(cfg); err != nil {
				return fmt.Errorf("failed to write config file because %+v", err)
			}

			return nil
		},
	}
}
