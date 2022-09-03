package entity

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/knakayama/dv/test"
	"github.com/stretchr/testify/assert"
)

func setupVpcTest(_ *testing.T) func(_ *testing.T) {
	test.RemoveVpcs()

	return func(_ *testing.T) {
		test.RemoveVpcs()
	}
}

func TestVpcNewInvalidClient(t *testing.T) {
	vpc, err := NewVpc(&ec2.Client{})

	assert.NotNil(t, vpc)
	assert.NotNil(t, err)
}

func TestVpcNewVpcNoVpc(t *testing.T) {
	teardownTest := setupVpcTest(t)
	defer teardownTest(t)

	vpc, err := NewVpc(NewDefaultClient())

	assert.Nil(t, vpc.Id)
	assert.Nil(t, err)
}

func TestVpcNewVpcNoDefaultVpc(t *testing.T) {
	teardownTest := setupVpcTest(t)
	defer teardownTest(t)
	test.CreateVpcs()

	vpc, err := NewVpc(NewDefaultClient())

	assert.Nil(t, vpc.Id)
	assert.Nil(t, err)
}

func TestVpcNewVpcDefaultVpcExists(t *testing.T) {
	teardownTest := setupVpcTest(t)
	defer teardownTest(t)
	test.CreateDefaultVpc()

	vpc, err := NewVpc(NewDefaultClient())

	assert.NotNil(t, vpc.Id)
	assert.Nil(t, err)
}

func TestVpcRemoveNoVpc(t *testing.T) {
	teardownTest := setupVpcTest(t)
	defer teardownTest(t)

	vpc, _ := NewVpc(NewDefaultClient())
	err := vpc.Remove()

	assert.ErrorIs(t, err, ErrVpcNotFound)
}

func TestVpcRemoveInvalidClient(t *testing.T) {
	teardownTest := setupVpcTest(t)
	defer teardownTest(t)

	vpc, _ := NewVpc(NewDefaultClient())
	vpc.Client = &ec2.Client{}
	err := vpc.Remove()

	assert.NotNil(t, err)
}

func TestVpcRemoveVpcExists(t *testing.T) {
	teardownTest := setupVpcTest(t)
	test.CreateDefaultVpc()
	defer teardownTest(t)

	vpc, _ := NewVpc(NewDefaultClient())
	err := vpc.Remove()

	assert.Nil(t, err)
}
