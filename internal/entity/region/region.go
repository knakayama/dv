package region

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func List(client *ec2.Client) (*ec2.DescribeRegionsOutput, error) {
	return client.DescribeRegions(
		context.TODO(),
		&ec2.DescribeRegionsInput{},
	)
}
