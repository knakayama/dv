package test

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func RemoveVpcs(client *ec2.Client) {
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

func CreateVpcs(client *ec2.Client) {
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

func CreateDefaultVpc(client *ec2.Client) {
	_, err := client.CreateDefaultVpc(
		context.TODO(),
		&ec2.CreateDefaultVpcInput{},
	)
	if err != nil {
		panic(err)
	}
}
