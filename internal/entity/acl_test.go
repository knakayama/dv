package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAclRemoveNoVpc(t *testing.T) {
	vpc, _ := NewVpc(NewDefaultClient())
	err := vpc.NewAcl().Remove()

	assert.Nil(t, err)
}
