package entity

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type Subnet struct {
	vpc *Vpc
}

func (subnet *Subnet) ids() ([]*string, error) {
	var subnetIds []*string

	output, err := subnet.vpc.Client.DescribeSubnets(
		context.TODO(),
		&ec2.DescribeSubnetsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*subnet.vpc.Id},
				},
			},
		},
	)
	if err != nil {
		return []*string{}, err
	}

	for _, subnet := range output.Subnets {
		subnetIds = append(subnetIds, subnet.SubnetId)
	}

	return subnetIds, nil
}

func (subnet *Subnet) Remove() error {
	subnetIds, _ := subnet.ids()

	for _, subnetId := range subnetIds {
		//nolint:forbidigo
		fmt.Println(*subnetId)
		_, err := subnet.vpc.Client.DeleteSubnet(
			context.TODO(),
			&ec2.DeleteSubnetInput{
				SubnetId: subnetId,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (vpc *Vpc) NewSubnet() *Subnet {
	return &Subnet{
		vpc: vpc,
	}
}
