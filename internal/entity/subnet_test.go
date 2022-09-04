package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubnetRemoveNoVpc(t *testing.T) {
	vpc, _ := NewVpc(NewDefaultClient())
	err := vpc.NewSubnet().Remove()

	assert.Nil(t, err)
}
