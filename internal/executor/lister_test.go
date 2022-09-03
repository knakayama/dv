package executor

import (
	"testing"

	"github.com/knakayama/dv/internal/entity"
	"github.com/knakayama/dv/test"
	"github.com/stretchr/testify/assert"
)

func TestListerListVpcsNoDefaultVpc(t *testing.T) {
	_ = RemoveVpcs(true)
	err := ListVpcs()

	assert.Nil(t, err)
}

func TestListerListVpcsVpcsExist(t *testing.T) {
	test.CreateVpcs(entity.NewDefaultClient())
	err := ListVpcs()

	assert.Nil(t, err)
}

func TestListerListVpcsDefaultVpcExists(t *testing.T) {
	test.CreateDefaultVpc(entity.NewDefaultClient())
	err := ListVpcs()

	assert.Nil(t, err)
}
