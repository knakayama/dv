package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecurityGroupRemoveNoVpc(t *testing.T) {
	vpc, _ := NewVpc(NewDefaultClient())
	err := vpc.NewSecurityGroup().Remove()

	assert.Nil(t, err)
}
