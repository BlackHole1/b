package asar

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/BlackHole1/b/pkg/e"
)

func Header(p string) error {
	asarPath, err := Path(p)
	if err != nil {
		return err
	}

	file, err := os.Open(asarPath)
	if err != nil {
		return err
	}
	defer file.Close()

	sBuf := make([]byte, 16)
	if n, _ := file.ReadAt(sBuf, 0); n != 16 {
		return e.AsarMalformed
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
