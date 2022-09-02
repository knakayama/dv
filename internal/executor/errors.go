package executor

import (
	"errors"
)

var (
	errVpcNotFound   = errors.New("vpc not found")
	errUnknownRegion = errors.New("unknown region")
)
