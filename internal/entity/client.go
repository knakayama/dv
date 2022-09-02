package entity

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

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

func NewDefaultClient() *ec2.Client {
	region := os.Getenv("DEFAULT_REGION")

	return ec2.NewFromConfig(makeConfig(region))
}

func NewClient(region string) *ec2.Client {
	return ec2.NewFromConfig(makeConfig(region))
}
