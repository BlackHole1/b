package util

import (
	"os"
)

func ExistFile(p string) bool {
	fileInfo, err := os.Stat(p)
	if err != nil {
		return false
	}

	if fileInfo.IsDir() {
		return false
	}

	return true
}

func ExistDir(p string) bool {
	fileInfo, err := os.Stat(p)
	if err != nil {
		return false
	}

	if fileInfo.IsDir() {
		return true
	}

	return false
}
