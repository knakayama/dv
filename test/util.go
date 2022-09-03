package test

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
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

func CreateVpcs() {
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

func CreateDefaultVpc() {
	_, err := client.CreateDefaultVpc(
		context.TODO(),
		&ec2.CreateDefaultVpcInput{},
	)
	if err != nil {
		panic(err)
	}
}

func CreateAcls(vpcId *string) {
	for i := 0; i < 5; i++ {
		_, err := client.CreateNetworkAcl(
			context.TODO(),
			&ec2.CreateNetworkAclInput{
				VpcId: vpcId,
			},
		)
		if err != nil {
			panic(err)
		}
	}
}

func ListAcls(vpcId *string) []types.NetworkAcl {
	out, err := client.DescribeNetworkAcls(
		context.TODO(),
		&ec2.DescribeNetworkAclsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*vpcId},
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}

	return out.NetworkAcls
}
