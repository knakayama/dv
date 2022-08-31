package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func DeleteSubnets(client *ec2.Client, vpc types.Vpc) {
	output, err := client.DescribeSubnets(
		context.TODO(),
		&ec2.DescribeSubnetsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*vpc.VpcId},
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("Failed to list subnets, %v", err)
	}

	for _, subnet := range output.Subnets {
		//nolint:forbidigo
		fmt.Println(*subnet.SubnetId)
		_, err := client.DeleteSubnet(
			context.TODO(),
			&ec2.DeleteSubnetInput{
				SubnetId: subnet.SubnetId,
			},
		)
		if err != nil {
			log.Fatalf("Failed to delete a subnet, %v", err)
		}
	}
}
