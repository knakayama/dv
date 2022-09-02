package service

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func DeleteIgws(client *ec2.Client, vpc types.Vpc) {
	out, err := client.DescribeInternetGateways(
		context.TODO(),
		&ec2.DescribeInternetGatewaysInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("attachment.vpc-id"),
					Values: []string{*vpc.VpcId},
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("Failed to list igws, %v", err)
	}

	for _, igw := range out.InternetGateways {
		//nolint:forbidigo
		fmt.Println(*igw.InternetGatewayId)
		_, err := client.DeleteInternetGateway(
			context.TODO(),
			&ec2.DeleteInternetGatewayInput{
				InternetGatewayId: igw.InternetGatewayId,
			},
		)
		if err != nil {
			log.Fatalf("Failed to delete an igw, %v", err)
		}
	}
}
