package entity

import (
	"errors"
)

var (
	ErrVpcNotFound   = errors.New("vpc not found")
	ErrUnknownRegion = errors.New("unknown region")
)
