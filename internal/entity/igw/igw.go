package igw

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/knakayama/dv/internal/util"
)

func ids(client *ec2.Client, vpcIds []*string) ([]*string, error) {
	var igwIds []*string

	output, err := client.DescribeInternetGateways(
		context.TODO(),
		&ec2.DescribeInternetGatewaysInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("attachment.vpc-id"),
					Values: util.StrListFrom(vpcIds),
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

func Remove(client *ec2.Client, vpcIds []*string) error {
	igwIds, _ := ids(client, vpcIds)

	for _, igwId := range igwIds {
		//nolint:forbidigo
		fmt.Println(*igwId)
		_, err := client.DeleteInternetGateway(
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
