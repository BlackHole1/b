package main

import (
	"os"

	"go.uber.org/zap"

	"github.com/spf13/cobra"

	"github.com/BlackHole1/b/pkg/log"
)

var debug bool

var rootCmd = &cobra.Command{
	Use:           "b",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "", false, "debug mode, also you can set DEBUG environment variable")

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if debug {
			log.SetLevel(zap.DebugLevel)
		}
		return nil
	}
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			log.Error("failed to print help", zap.Error(err))
			os.Exit(1)
		}
	}
}
