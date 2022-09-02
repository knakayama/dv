package entity

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type RouteTable struct {
	vpc *Vpc
}

func (r *RouteTable) ids() ([]*string, error) {
	var routeTableIds []*string

	out, err := r.vpc.Client.DescribeRouteTables(
		context.TODO(),
		&ec2.DescribeRouteTablesInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*r.vpc.Id},
				},
			},
		},
	)
	if err != nil {
		return []*string{}, err
	}

	for _, routeTable := range out.RouteTables {
		routeTableIds = append(routeTableIds, routeTable.RouteTableId)
	}

	return routeTableIds, nil
}

func (r *RouteTable) Remove() error {
	routeTableIds, _ := r.ids()

	for _, routeTableId := range routeTableIds {
		//nolint:forbidigo
		fmt.Println(*routeTableId)
		_, err := r.vpc.Client.DeleteRouteTable(
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

func (v *Vpc) NewRouteTable() *RouteTable {
	return &RouteTable{
		vpc: v,
	}
}
