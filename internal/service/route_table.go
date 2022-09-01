package service

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func DeleteRouteTables(client *ec2.Client, vpc types.Vpc) {
	output, err := client.DescribeRouteTables(
		context.TODO(),
		&ec2.DescribeRouteTablesInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*vpc.VpcId},
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("Failed to list route tables, %v", err)
	}

	for _, routeTable := range output.RouteTables {
		//nolint:forbidigo
		fmt.Println(*routeTable.RouteTableId)
		_, err := client.DeleteRouteTable(
			context.TODO(),
			&ec2.DeleteRouteTableInput{
				RouteTableId: routeTable.RouteTableId,
			},
		)
		if err != nil {
			log.Fatalf("Failed to delete a route table, %v", err)
		}
	}
}
