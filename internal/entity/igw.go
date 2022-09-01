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

func (igw *Igw) ids() ([]*string, error) {
	var igwIds []*string

	output, err := igw.vpc.Client.DescribeInternetGateways(
		context.TODO(),
		&ec2.DescribeInternetGatewaysInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("attachment.vpc-id"),
					Values: []string{*igw.vpc.Id},
				},
			},
		},
	)
	if err != nil {
		return []*string{}, err
	}

	for _, igw := range output.InternetGateways {
		igwIds = append(igwIds, igw.InternetGatewayId)
	}

	return igwIds, nil
}

func (igw *Igw) Remove() error {
	igwIds, _ := igw.ids()

	for _, igwId := range igwIds {
		//nolint:forbidigo
		fmt.Println(*igwId)
		_, err := igw.vpc.Client.DeleteInternetGateway(
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

func (vpc *Vpc) NewIgw() *Igw {
	return &Igw{
		vpc: vpc,
	}
}
