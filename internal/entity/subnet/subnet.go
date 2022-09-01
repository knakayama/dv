package subnet

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/knakayama/dv/internal/util"
)

func ids(client *ec2.Client, vpcIds []*string) ([]*string, error) {
	var subnetIds []*string

	output, err := client.DescribeSubnets(
		context.TODO(),
		&ec2.DescribeSubnetsInput{
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

	for _, subnet := range output.Subnets {
		subnetIds = append(subnetIds, subnet.SubnetId)
	}

	return subnetIds, nil
}

func Remove(client *ec2.Client, vpcIds []*string) error {
	subnetIds, _ := ids(client, vpcIds)

	for _, subnetId := range subnetIds {
		//nolint:forbidigo
		fmt.Println(*subnetId)
		_, err := client.DeleteSubnet(
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
