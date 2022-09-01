package entity

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type Region struct {
	client *ec2.Client
}

func (region *Region) List() (*ec2.DescribeRegionsOutput, error) {
	return region.client.DescribeRegions(
		context.TODO(),
		&ec2.DescribeRegionsInput{},
	)
}

func NewRegion(client *ec2.Client) *Region {
	return &Region{
		client: client,
	}
}
