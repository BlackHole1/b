/*
 * SPDX-FileCopyrightText: 2024 Kevin Cui <bh@bugs.cc>
 * SPDX-License-Identifier: MPL-2.0
 */

package registry

import "github.com/spf13/cobra"

type CliCommand struct {
	Command *cobra.Command
	Parent  *cobra.Command
}

var (
	Commands []CliCommand
)

func AddCommand(parent *cobra.Command, cmd *cobra.Command) {
	Commands = append(Commands, CliCommand{
		Command: cmd,
		Parent:  parent,
	})
}
