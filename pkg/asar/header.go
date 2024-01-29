/*
 * SPDX-FileCopyrightText: 2024 Kevin Cui <bh@bugs.cc>
 * SPDX-License-Identifier: MPL-2.0
 */

package asar

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/BlackHole1/b/cmd/b/registry"
)

func Header(p string) error {
	file, err := os.Open(p)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			registry.SetExitCode(1)
		}
		return err
	}
	defer file.Close()

	sBuf := make([]byte, 16)
	if n, _ := file.ReadAt(sBuf, 0); n != 16 {
		return errors.New("asar malformed archive")
	}

	size := binary.LittleEndian.Uint32(sBuf[12:16])

	sr := io.NewSectionReader(file, 16, int64(size))

	hBuf := make([]byte, size)
	if _, err = sr.Read(hBuf); err != nil {
		return err
	}

	fmt.Print(string(hBuf))

	return nil
}
