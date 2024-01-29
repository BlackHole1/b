/*
 * SPDX-FileCopyrightText: 2024 Kevin Cui <bh@bugs.cc>
 * SPDX-License-Identifier: MPL-2.0
 */
// From: https://github.com/containers/podman/blob/d7bf1385d8894f7ef1d002917492489e080a5e0b/cmd/podman/validate/args.go

package validate

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// SubCommandExists returns an error if no sub command is provided
func SubCommandExists(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		fmt.Println(args, cmd.IsAvailableCommand())
		suggestions := cmd.SuggestionsFor(args[0])
		if len(suggestions) == 0 {
			return fmt.Errorf("unrecognized command `%[1]s %[2]s`\nTry '%[1]s --help' for more information", cmd.CommandPath(), args[0])
		}
		return fmt.Errorf("unrecognized command `%[1]s %[2]s`\n\nDid you mean this?\n\t%[3]s\n\nTry '%[1]s --help' for more information", cmd.CommandPath(), args[0], strings.Join(suggestions, "\n\t"))
	}

	_ = cmd.Help()
	return fmt.Errorf("missing command '%[1]s COMMAND'", cmd.CommandPath())
}

// NoArgs returns an error if any args are included.
func NoArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("`%s` takes no arguments", cmd.CommandPath())
	}
	return nil
}
