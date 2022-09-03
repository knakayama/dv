package entity

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/stretchr/testify/assert"
)

func TestAclIdsNoVpc(t *testing.T) {
	client := NewDefaultClient()
	teardown := setup(t, client)
	defer teardown(t)

	vpc, _ := NewVpc(client)
	aclIds, err := vpc.NewAcl().ids()

	assert.Empty(t, aclIds)
	assert.ErrorIs(t, err, ErrVpcNotFound)
}

func TestAclRemoveNoVpc(t *testing.T) {
	client := NewDefaultClient()
	teardown := setup(t, client)
	defer teardown(t)

	vpc, _ := NewVpc(client)
	err := vpc.NewAcl().Remove()

	assert.NotNil(t, err)
}

// Default VPC always has network components such as ACL.
func TestAclRemoveAclsExist(t *testing.T) {
	client := NewDefaultClient()
	teardown := setupVpc(t, client)
	defer teardown(t)

	vpc, _ := NewVpc(client)
	createAcls(vpc)
	err := vpc.NewAcl().Remove()
	acls := listAcls(vpc)

	assert.Nil(t, err)
	assert.Empty(t, acls)
}

func createAcls(vpc *Vpc) {
	for i := 0; i < 5; i++ {
		_, err := vpc.Client.CreateNetworkAcl(
			context.TODO(),
			&ec2.CreateNetworkAclInput{
				VpcId: vpc.Id,
			},
		)
		if err != nil {
			panic(err)
		}
	}
}

func listAcls(vpc *Vpc) []types.NetworkAcl {
	out, err := vpc.Client.DescribeNetworkAcls(
		context.TODO(),
		&ec2.DescribeNetworkAclsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*vpc.Id},
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}

	return out.NetworkAcls
}
