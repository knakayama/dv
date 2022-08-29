package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func makeClient() *ec2.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalf("Failed to load SDK configuration, %v", err)
	}

	return ec2.NewFromConfig(cfg)
}

func listRegions(client *ec2.Client) []types.Region {
	output, err := client.DescribeRegions(
		context.TODO(),
		&ec2.DescribeRegionsInput{},
	)

	if err != nil {
		log.Fatalf("Failed to list regions, %v", err)
	}

	return output.Regions
}

func makeClients() []*ec2.Client {
	var clients []*ec2.Client

	for _, region := range listRegions(makeClient()) {
		cfg, err := config.LoadDefaultConfig(
			context.TODO(),
			config.WithRegion(*region.RegionName),
		)
		if err != nil {
			log.Fatalf("Failed to load SDK configuration, %v", err)
		}
		clients = append(clients, ec2.NewFromConfig(cfg))
	}

	return clients
}

func listDefaultVpcs(client *ec2.Client) []types.Vpc {
	var vpcs []types.Vpc

	output, err := client.DescribeVpcs(
		context.TODO(),
		&ec2.DescribeVpcsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("isDefault"),
					Values: []string{"true"},
				},
			},
		})

	if err != nil {
		log.Fatalf("Failed to list default vpcs, %v", err)
	}

	for _, vpc := range output.Vpcs {
		if *vpc.IsDefault {
			vpcs = append(vpcs, vpc)
		}
	}

	return vpcs
}

func main() {
	for _, client := range makeClients() {
		for _, vpc := range listDefaultVpcs(client) {
			fmt.Println(*vpc.VpcId)
		}
	}
}
