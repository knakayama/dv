package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func DeleteAcls(client *ec2.Client, vpc types.Vpc) {
	output, err := client.DescribeNetworkAcls(
		context.TODO(),
		&ec2.DescribeNetworkAclsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*vpc.VpcId},
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("Failed to list network acls, %v", err)
	}

	for _, acl := range output.NetworkAcls {
		//nolint:forbidigo
		fmt.Println(*acl.NetworkAclId)
		_, err := client.DeleteNetworkAcl(
			context.TODO(),
			&ec2.DeleteNetworkAclInput{
				NetworkAclId: acl.NetworkAclId,
			},
		)
		if err != nil {
			log.Fatalf("Failed to delete a network acl, %v", err)
		}
	}
}
