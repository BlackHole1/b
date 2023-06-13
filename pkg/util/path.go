package util

import (
	"path"
	"path/filepath"
)

func PathToAbs(p string) (string, error) {
	if path.IsAbs(p) {
		return p, nil
	}

	absPath, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}

	return absPath, nil
}
