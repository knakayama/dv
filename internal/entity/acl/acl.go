package acl

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/knakayama/dv/internal/util"
)

func ids(client *ec2.Client, vpcIds []*string) ([]*string, error) {
	var aclIds []*string

	output, err := client.DescribeNetworkAcls(
		context.TODO(),
		&ec2.DescribeNetworkAclsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: util.StrListFrom(vpcIds),
				},
			},
		},
	)
	if err != nil {
		return []*string{}, err
	}

	for _, acl := range output.NetworkAcls {
		aclIds = append(aclIds, acl.NetworkAclId)
	}

	return aclIds, nil
}

func Remove(client *ec2.Client, vpcIds []*string) error {
	aclIds, _ := ids(client, vpcIds)

	for _, aclId := range aclIds {
		//nolint:forbidigo
		fmt.Println(*aclId)
		_, err := client.DeleteNetworkAcl(
			context.TODO(),
			&ec2.DeleteNetworkAclInput{
				NetworkAclId: aclId,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}
