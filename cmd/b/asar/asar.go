/*
 * SPDX-FileCopyrightText: 2024 Kevin Cui <bh@bugs.cc>
 * SPDX-License-Identifier: MPL-2.0
 */

package asar

import (
	"github.com/BlackHole1/b/cmd/b/registry"
	"github.com/BlackHole1/b/cmd/b/validate"
	"github.com/spf13/cobra"
)

var (
	asarCmd = &cobra.Command{
		Use:   "asar",
		Short: "Electron asar",
		Long:  "Helper tool for electron asar",
		RunE:  validate.SubCommandExists,
	}
)

func init() {
	registry.AddCommand(nil, asarCmd)
}
