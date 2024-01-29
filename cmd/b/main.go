/*
 * SPDX-FileCopyrightText: 2024 Kevin Cui <bh@bugs.cc>
 * SPDX-License-Identifier: MPL-2.0
 */

package main

import (
	"fmt"

	_ "github.com/BlackHole1/b/cmd/b/asar"
	"github.com/BlackHole1/b/pkg/define"
)

import (
	"os"

	"github.com/BlackHole1/b/cmd/b/registry"
	"github.com/BlackHole1/b/internal/version"
	"github.com/BlackHole1/b/pkg/log"
)

func main() {
	defer log.Sync()

	rootCmd.Version = version.Version

	for _, command := range registry.Commands {
		parent := rootCmd
		if command.Parent != nil {
			parent = command.Parent
		}

		parent.AddCommand(command.Command)
	}

	if err := rootCmd.Execute(); err != nil {
		if registry.GetExitCode() == 0 {
			registry.SetExitCode(define.ExecErrorCodeGeneric)
		}

		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(registry.GetExitCode())
	}
}
