package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIgwRemoveNoVpc(t *testing.T) {
	vpc, _ := NewVpc(NewDefaultClient())
	err := vpc.NewIgw().Remove()

	assert.Nil(t, err)
}
