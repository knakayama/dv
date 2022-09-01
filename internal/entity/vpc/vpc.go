package vpc

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func Remove(client *ec2.Client, vpcIds []*string) error {
	for _, vpcId := range vpcIds {
		_, err := client.DeleteVpc(
			context.TODO(),
			&ec2.DeleteVpcInput{
				VpcId: vpcId,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func Ids(client *ec2.Client) ([]*string, error) {
	var vpcIds []*string

	output, err := client.DescribeVpcs(
		context.TODO(),
		&ec2.DescribeVpcsInput{})
	if err != nil {
		return []*string{}, err
	}

	for _, vpc := range output.Vpcs {
		if *vpc.IsDefault {
			vpcIds = append(vpcIds, vpc.VpcId)
		}
	}

	return vpcIds, nil
}
