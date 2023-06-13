package main

import (
	"os"

	"go.uber.org/zap"

	"github.com/spf13/cobra"

	"github.com/BlackHole1/b/pkg/asar"
	"github.com/BlackHole1/b/pkg/log"
)

var asarCmd = &cobra.Command{
	Use:   "asar",
	Short: "Electron asar",
	Long:  "Helper tool for electron asar",
}

var headerCmd = &cobra.Command{
	Use:   "header",
	Short: "Print asar header",
	Long:  "Print asar header json",
}

func init() {
	initHeaderCmd()

	asarCmd.AddCommand(headerCmd)
}

func initHeaderCmd() {
	const (
		flagFile = "file"
	)
	var (
		file string
	)

	headerCmd.Flags().StringVarP(&file, flagFile, "f", "", "asar file path")
	if err := headerCmd.MarkFlagFilename(flagFile, "asar"); err != nil {
		log.Error("failed to mark flag as filename", zap.Error(err))
		os.Exit(1)
	}

	headerCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return asar.Header(file)
	}
}
