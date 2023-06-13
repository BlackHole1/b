package e

import (
	"errors"
)

var AsarNotFound = errors.New("asar not found")
var NotAsarFile = errors.New("not asar file")
var AsarMalformed = errors.New("asar malformed archive")
