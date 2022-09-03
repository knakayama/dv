package entity

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/stretchr/testify/assert"
)

func setup(_ *testing.T, client *ec2.Client) func(_ *testing.T) {
	removeVpcs(client)

	return func(_ *testing.T) {
		removeVpcs(client)
	}
}

func setupVpc(_ *testing.T, client *ec2.Client) func(_ *testing.T) {
	removeVpcs(client)
	createDefaultVpc(client)

	return func(_ *testing.T) {
		removeVpcs(client)
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
	createVpcs(client)

	vpc, err := NewVpc(client)

	assert.Nil(t, vpc.Id)
	assert.Nil(t, err)
}

func TestVpcNewVpcDefaultVpcExists(t *testing.T) {
	client := NewDefaultClient()
	teardown := setup(t, client)
	defer teardown(t)
	createDefaultVpc(client)

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
	createDefaultVpc(client)
	defer teardown(t)

	vpc, _ := NewVpc(client)
	err := vpc.Remove()

	assert.Nil(t, err)
}

func removeVpcs(client *ec2.Client) {
	out, err := client.DescribeVpcs(
		context.TODO(),
		&ec2.DescribeVpcsInput{},
	)
	if err != nil {
		panic(err)
	}

	for _, vpc := range out.Vpcs {
		_, err := client.DeleteVpc(
			context.TODO(),
			&ec2.DeleteVpcInput{
				VpcId: vpc.VpcId,
			},
		)
		if err != nil {
			panic(err)
		}
	}
}

func createVpcs(client *ec2.Client) {
	for i := 0; i < 5; i++ {
		_, err := client.CreateVpc(
			context.TODO(),
			&ec2.CreateVpcInput{
				CidrBlock: aws.String("192.168.0.0/16"),
			},
		)
		if err != nil {
			panic(err)
		}
	}
}

func createDefaultVpc(client *ec2.Client) {
	_, err := client.CreateDefaultVpc(
		context.TODO(),
		&ec2.CreateDefaultVpcInput{},
	)
	if err != nil {
		panic(err)
	}
}
