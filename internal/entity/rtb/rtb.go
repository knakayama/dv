package rtb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/knakayama/dv/internal/util"
)

func ids(client *ec2.Client, vpcIds []*string) ([]*string, error) {
	var routeTableIds []*string

	output, err := client.DescribeRouteTables(
		context.TODO(),
		&ec2.DescribeRouteTablesInput{
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

	for _, routeTable := range output.RouteTables {
		routeTableIds = append(routeTableIds, routeTable.RouteTableId)
	}

	return routeTableIds, nil
}

func Remove(client *ec2.Client, vpcIds []*string) error {
	routeTableIds, _ := ids(client, vpcIds)

	for _, routeTableId := range routeTableIds {
		//nolint:forbidigo
		fmt.Println(*routeTableId)
		_, err := client.DeleteRouteTable(
			context.TODO(),
			&ec2.DeleteRouteTableInput{
				RouteTableId: routeTableId,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}
