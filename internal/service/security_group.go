package service

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func DeleteSecurityGroups(client *ec2.Client, vpc types.Vpc) {
	output, err := client.DescribeSecurityGroups(
		context.TODO(),
		&ec2.DescribeSecurityGroupsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*vpc.VpcId},
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("Failed to list security groups, %v", err)
	}

	for _, securityGroup := range output.SecurityGroups {
		_, err := client.DeleteSecurityGroup(
			context.TODO(),
			&ec2.DeleteSecurityGroupInput{
				GroupId: securityGroup.GroupId,
			},
		)
		if err != nil {
			log.Fatalf("Failed to delete a security groups, %v", err)
		}
	}
}
