/*
 * SPDX-FileCopyrightText: 2024 Kevin Cui <bh@bugs.cc>
 * SPDX-License-Identifier: MPL-2.0
 */

package asar

import (
	"errors"

	"github.com/BlackHole1/b/cmd/b/completion"
	"github.com/BlackHole1/b/cmd/b/registry"
	"github.com/BlackHole1/b/cmd/b/validate"
	"github.com/BlackHole1/b/pkg/asar"
	"github.com/spf13/cobra"
)

var (
	inFile string
)

var (
	headerCmd = &cobra.Command{
		Use:               "header",
		Short:             "Print asar header",
		Long:              "Print asar header json",
		Args:              validate.NoArgs,
		RunE:              printJSON,
		ValidArgsFunction: completion.AutocompleteNone,
	}
)

func init() {
	registry.AddCommand(asarCmd, headerCmd)

	const fileFlagName = "file"

	headerCmd.Flags().StringVarP(&inFile, fileFlagName, "f", "", "asar file path")
	_ = headerCmd.MarkFlagFilename(fileFlagName, "asar")
}

func printJSON(cmd *cobra.Command, args []string) error {
	if len(inFile) == 0 {
		return errors.New("cannot read from terminal, use command-line redirection or the --file flag")
	}

	return asar.Header(inFile)
}
