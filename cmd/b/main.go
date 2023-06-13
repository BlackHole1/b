package main

import (
	"os"

	"go.uber.org/zap"

	"github.com/BlackHole1/b/pkg/log"
)

func main() {
	defer log.Sync()
	rootCmd.AddCommand(asarCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Error("failed to execute command", zap.Error(err))
		os.Exit(1)
	}
}
