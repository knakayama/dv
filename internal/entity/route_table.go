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
		// We can't delete Main route table.
		if *routeTable.Associations[0].Main {
			//nolint:forbidigo
			fmt.Printf("%s is Main route table, skipped...\n", *routeTable.RouteTableId)
			continue
		}
		routeTableIds = append(routeTableIds, routeTable.RouteTableId)
	}

	return routeTableIds, nil
}

func (r *RouteTable) Remove() error {
	if r.vpc.Id == nil {
		return nil
	}

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
