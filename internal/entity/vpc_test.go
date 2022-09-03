package entity

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/knakayama/dv/test"
	"github.com/stretchr/testify/assert"
)

func setup(_ *testing.T, client *ec2.Client) func(_ *testing.T) {
	test.RemoveVpcs(client)

	return func(_ *testing.T) {
		test.RemoveVpcs(client)
	}
}

func TestVpcNewInvalidClient(t *testing.T) {
	vpc, err := NewVpc(&ec2.Client{})

	assert.NotNil(t, vpc)
	assert.NotNil(t, err)
}

func TestVpcNewVpcNoVpc(t *testing.T) {
	client := NewDefaultClient()
	teardown := setup(t, client)
	defer teardown(t)

	vpc, err := NewVpc(client)

	assert.Nil(t, vpc.Id)
	assert.Nil(t, err)
}

func TestVpcNewVpcNoDefaultVpc(t *testing.T) {
	client := NewDefaultClient()
	teardown := setup(t, client)
	defer teardown(t)
	test.CreateVpcs(client)

	vpc, err := NewVpc(client)

	assert.Nil(t, vpc.Id)
	assert.Nil(t, err)
}

func TestVpcNewVpcDefaultVpcExists(t *testing.T) {
	client := NewDefaultClient()
	teardown := setup(t, client)
	defer teardown(t)
	test.CreateDefaultVpc(client)

	vpc, err := NewVpc(client)

	assert.NotNil(t, vpc.Id)
	assert.Nil(t, err)
}

func TestVpcRemoveNoVpc(t *testing.T) {
	client := NewDefaultClient()
	teardown := setup(t, client)
	defer teardown(t)

	vpc, _ := NewVpc(client)
	err := vpc.Remove()

	assert.ErrorIs(t, err, ErrVpcNotFound)
}

func TestVpcRemoveInvalidClient(t *testing.T) {
	client := NewDefaultClient()
	teardown := setup(t, client)
	defer teardown(t)

	vpc, _ := NewVpc(client)
	vpc.Client = &ec2.Client{}
	err := vpc.Remove()

	assert.NotNil(t, err)
}

func TestVpcRemoveVpcExists(t *testing.T) {
	client := NewDefaultClient()
	teardown := setup(t, client)
	test.CreateDefaultVpc(client)
	defer teardown(t)

	vpc, _ := NewVpc(client)
	err := vpc.Remove()

	assert.Nil(t, err)
}
