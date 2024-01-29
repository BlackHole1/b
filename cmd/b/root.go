/*
 * SPDX-FileCopyrightText: 2024 Kevin Cui <bh@bugs.cc>
 * SPDX-License-Identifier: MPL-2.0
 */

package main

import (
	"github.com/BlackHole1/b/cmd/b/validate"
	"go.uber.org/zap"

	"github.com/spf13/cobra"

	"github.com/BlackHole1/b/pkg/log"
)

var (
	inDebug bool
)

var rootCmd = &cobra.Command{
	Use:              "b",
	SilenceUsage:     true,
	SilenceErrors:    true,
	RunE:             validate.SubCommandExists,
	PersistentPreRun: setLogLevel,
}

func init() {
	const debugFlagName = "debug"
	rootCmd.PersistentFlags().BoolVarP(&inDebug, debugFlagName, "", false, "debug mode, also you can set DEBUG environment variable")
}

func setLogLevel(cmd *cobra.Command, args []string) {
	if inDebug {
		log.SetLevel(zap.DebugLevel)
	}
}
