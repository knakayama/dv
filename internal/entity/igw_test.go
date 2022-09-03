package entity

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/stretchr/testify/assert"
)

func TestIgwIdsNoVpc(t *testing.T) {
	client := NewDefaultClient()
	teardown := setup(t, client)
	defer teardown(t)

	vpc, _ := NewVpc(client)
	igwIds, err := vpc.NewIgw().ids()

	assert.Empty(t, igwIds)
	assert.ErrorIs(t, err, ErrVpcNotFound)
}

func TestIgwIdsVpcExists(t *testing.T) {
	client := NewDefaultClient()
	teardown := setupVpc(t, client)
	defer teardown(t)

	vpc, _ := NewVpc(client)
	createIgw(vpc)
	igwIds, err := vpc.NewIgw().ids()

	assert.NotEmpty(t, igwIds)
	assert.Nil(t, err)
}

func TestIgwRemoveNoVpc(t *testing.T) {
	client := NewDefaultClient()
	teardown := setup(t, client)
	defer teardown(t)

	vpc, _ := NewVpc(client)
	err := vpc.NewIgw().Remove()

	assert.Error(t, err)
}

// TODO: Localstack doesn't support DescribeInternetGateways yet.
//func TestIgwRemoveIgwExists(t *testing.T) {
//	client := NewDefaultClient()
//	teardown := setupVpc(t, client)
//	defer teardown(t)
//
//	vpc, _ := NewVpc(client)
//	err := vpc.NewIgw().Remove()
//	igws := listIgw(vpc)
//
//	assert.Nil(t, err)
//	assert.Empty(t, igws)
//}
//
//func listIgw(vpc *Vpc) []types.InternetGateway {
//	out, err := vpc.Client.DescribeInternetGateways(
//		context.TODO(),
//		&ec2.DescribeInternetGatewaysInput{
//			Filters: []types.Filter{
//				{
//					Name:   aws.String("vpc-id"),
//					Values: []string{*vpc.Id},
//				},
//			},
//		},
//	)
//	if err != nil {
//		panic(err)
//	}
//
//	return out.InternetGateways
//}

func createIgw(vpc *Vpc) {
	out, err := vpc.Client.CreateInternetGateway(
		context.TODO(),
		&ec2.CreateInternetGatewayInput{},
	)
	if err != nil {
		panic(err)
	}

	_, err = vpc.Client.AttachInternetGateway(
		context.TODO(),
		&ec2.AttachInternetGatewayInput{
			VpcId:             vpc.Id,
			InternetGatewayId: out.InternetGateway.InternetGatewayId,
		},
	)
	if err != nil {
		panic(err)
	}
}
