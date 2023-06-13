package asar

import (
	"os"
	"path"

	"github.com/BlackHole1/b/pkg/e"
	"github.com/BlackHole1/b/pkg/util"
)

func Path(p string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if p != "" {
		abs, err := util.PathToAbs(p)
		if err != nil {
			return "", err
		}

		if util.ExistFile(abs) {
			if path.Ext(abs) != ".asar" {
				return "", e.NotAsarFile
			}
			return abs, nil
		}

		return "", e.AsarNotFound
	}

	for _, f := range asarFilenames {
		if util.ExistFile(path.Join(dir, f)) {
			p = path.Join(dir, f)
			break
		}
	}

	if p == "" {
		return "", e.AsarNotFound
	}

	return p, nil
}
