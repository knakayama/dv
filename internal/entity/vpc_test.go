package entity

import (
	"testing"

	"github.com/knakayama/dv/test"
	"github.com/stretchr/testify/assert"
)

func setupTest(_ *testing.T) func(_ *testing.T) {
	test.RemoveVpcs()

	return func(_ *testing.T) {
		test.RemoveVpcs()
	}
}

func TestNewVpcNoVpc(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	vpc, _ := NewVpc(NewDefaultClient())

	assert.Nil(t, vpc.Id)
}

func TestNewVpcNoDefaultVpc(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	test.MakeVpcs()

	vpc, _ := NewVpc(NewDefaultClient())

	assert.Nil(t, vpc.Id)
}

func TestNewVpcDefaultVpcFound(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	test.MakeDefaultVpc()

	vpc, _ := NewVpc(NewDefaultClient())

	assert.NotNil(t, vpc.Id)
}
