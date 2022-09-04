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

func (s *Subnet) ids() ([]*string, error) {
	var subnetIds []*string

	out, err := s.vpc.Client.DescribeSubnets(
		context.TODO(),
		&ec2.DescribeSubnetsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*s.vpc.Id},
				},
			},
		},
	)
	if err != nil {
		return []*string{}, err
	}

	for _, subnet := range out.Subnets {
		subnetIds = append(subnetIds, subnet.SubnetId)
	}

	return subnetIds, nil
}

func (s *Subnet) Remove() error {
	if s.vpc.Id == nil {
		return nil
	}

	subnetIds, _ := s.ids()

	for _, subnetId := range subnetIds {
		//nolint:forbidigo
		fmt.Println(*subnetId)
		_, err := s.vpc.Client.DeleteSubnet(
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

func (v *Vpc) NewSubnet() *Subnet {
	return &Subnet{
		vpc: v,
	}
}
