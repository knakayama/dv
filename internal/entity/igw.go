package entity

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type Igw struct {
	vpc *Vpc
}

func (i *Igw) ids() ([]*string, error) {
	var igwIds []*string

	out, err := i.vpc.Client.DescribeInternetGateways(
		context.TODO(),
		&ec2.DescribeInternetGatewaysInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("attachment.vpc-id"),
					Values: []string{*i.vpc.Id},
				},
			},
		},
	)
	if err != nil {
		return []*string{}, err
	}

	for _, igw := range out.InternetGateways {
		igwIds = append(igwIds, igw.InternetGatewayId)
	}

	return igwIds, nil
}

func (i *Igw) Remove() error {
	if i.vpc.Id == nil {
		return nil
	}

	igwIds, _ := i.ids()

	for _, igwId := range igwIds {
		//nolint:forbidigo
		fmt.Println(*igwId)
		_, err := i.vpc.Client.DeleteInternetGateway(
			context.TODO(),
			&ec2.DeleteInternetGatewayInput{
				InternetGatewayId: igwId,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *Vpc) NewIgw() *Igw {
	return &Igw{
		vpc: v,
	}
}
