package entity

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/knakayama/dv/test"
	"github.com/stretchr/testify/assert"
)

func setupTest(_ *testing.T) func(_ *testing.T) {
	test.RemoveVpcs()

	return func(_ *testing.T) {
		test.RemoveVpcs()
	}
}

func TestNewClientError(t *testing.T) {
	vpc, err := NewVpc(&ec2.Client{})

	assert.NotNil(t, vpc)
	assert.NotNil(t, err)
}

func TestNewVpcNoVpc(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	vpc, err := NewVpc(NewDefaultClient())

	assert.Nil(t, vpc.Id)
	assert.Nil(t, err)
}

func TestNewVpcNoDefaultVpc(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	test.MakeVpcs()

	vpc, err := NewVpc(NewDefaultClient())

	assert.Nil(t, vpc.Id)
	assert.Nil(t, err)
}

func TestNewVpcDefaultVpcFound(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	test.MakeDefaultVpc()

	vpc, err := NewVpc(NewDefaultClient())

	assert.NotNil(t, vpc.Id)
	assert.Nil(t, err)
}

func TestRemoveNoVpc(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	vpc, _ := NewVpc(NewDefaultClient())
	err := vpc.Remove()

	assert.ErrorIs(t, err, ErrVpcNotFound)
}

func TestRemoveInvalidClient(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	vpc, _ := NewVpc(NewDefaultClient())
	vpc.Client = &ec2.Client{}
	err := vpc.Remove()

	assert.NotNil(t, err)
}

func TestRemoveVpcExists(t *testing.T) {
	teardownTest := setupTest(t)
	test.MakeDefaultVpc()
	defer teardownTest(t)

	vpc, _ := NewVpc(NewDefaultClient())
	err := vpc.Remove()

	assert.Nil(t, err)
}
