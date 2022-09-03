package test

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var client *ec2.Client

func init() {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(os.Getenv("DEFAULT_REGION")),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					PartitionID:   "aws",
					URL:           os.Getenv("AWS_ENDPOINT"),
					SigningRegion: os.Getenv("DEFAULT_REGION"),
				}, nil
			},
			),
		))
	if err != nil {
		panic(err)
	}

	client = ec2.NewFromConfig(cfg)
}

func RemoveVpcs() {
	out, err := client.DescribeVpcs(
		context.TODO(),
		&ec2.DescribeVpcsInput{},
	)
	if err != nil {
		panic(err)
	}

	for _, vpc := range out.Vpcs {
		_, err := client.DeleteVpc(
			context.TODO(),
			&ec2.DeleteVpcInput{
				VpcId: vpc.VpcId,
			},
		)
		if err != nil {
			panic(err)
		}
	}
}

func MakeVpcs() {
	for i := 0; i < 5; i++ {
		_, err := client.CreateVpc(
			context.TODO(),
			&ec2.CreateVpcInput{
				CidrBlock: aws.String("192.168.0.0/16"),
			},
		)
		if err != nil {
			panic(err)
		}
	}
}

func MakeDefaultVpc() {
	_, err := client.CreateDefaultVpc(
		context.TODO(),
		&ec2.CreateDefaultVpcInput{},
	)
	if err != nil {
		panic(err)
	}
}
