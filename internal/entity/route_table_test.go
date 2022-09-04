package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouteTableRemoveNoVpc(t *testing.T) {
	vpc, _ := NewVpc(NewDefaultClient())
	err := vpc.NewRouteTable().Remove()

	assert.Nil(t, err)
}
