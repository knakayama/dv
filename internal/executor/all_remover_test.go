package executor

import (
	"testing"

	"github.com/knakayama/dv/internal/entity"
	"github.com/knakayama/dv/test"
	"github.com/stretchr/testify/assert"
)

func TestAllRemoverRemoveVpcsNoVpc(t *testing.T) {
	test.RemoveVpcs(entity.NewDefaultClient())

	err := RemoveVpcs(true)

	assert.Nil(t, err)
}

func TestAllRemoverRemoveVpcsVpcsExist(t *testing.T) {
	client := entity.NewDefaultClient()
	test.RemoveVpcs(client)
	test.CreateDefaultVpc(client)

	err := RemoveVpcs(true)

	assert.Nil(t, err)
}
