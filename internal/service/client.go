package service

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func listRegions(client *ec2.Client) []types.Region {
	out, err := client.DescribeRegions(
		context.TODO(),
		&ec2.DescribeRegionsInput{},
	)
	if err != nil {
		log.Fatalf("Failed to list regions, %v", err)
	}

	return out.Regions
}

func makeConfig(region string) aws.Config {
	awsEndpoint := os.Getenv("AWS_ENDPOINT")

	endpointResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: region,
			}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithEndpointResolverWithOptions(endpointResolver))
	if err != nil {
		log.Fatalf("Failed to load SDK configuration, %v", err)
	}

	return cfg
}

func makeClient() *ec2.Client {
	region := os.Getenv("DEFAULT_REGION")

	return ec2.NewFromConfig(makeConfig(region))
}

func MakeClients() []*ec2.Client {
	var clients []*ec2.Client

	for _, region := range listRegions(makeClient()) {
		clients = append(clients, ec2.NewFromConfig(makeConfig(*region.RegionName)))
	}

	return clients
}
