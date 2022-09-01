package service

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/knakayama/dv/internal/entity/acl"
	"github.com/knakayama/dv/internal/entity/client"
	"github.com/knakayama/dv/internal/entity/igw"
	"github.com/knakayama/dv/internal/entity/rtb"
	"github.com/knakayama/dv/internal/entity/sg"
	"github.com/knakayama/dv/internal/entity/subnet"
	"github.com/knakayama/dv/internal/entity/vpc"
)

func ListDefaultVpcs(client *ec2.Client) []types.Vpc {
	var vpcs []types.Vpc

	output, err := client.DescribeVpcs(
		context.TODO(),
		&ec2.DescribeVpcsInput{})
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

func DeleteVpc(client *ec2.Client, vpc types.Vpc) {
	_, err := client.DeleteVpc(
		context.TODO(),
		&ec2.DeleteVpcInput{
			VpcId: vpc.VpcId,
		},
	)
	if err != nil {
		log.Fatalf("Failed to delete a vpc, %v", err)
	}
}

func RemoveVpc() error {
	c := client.New()

	vpcIds, err := vpc.Ids(c)
	if err != nil {
		return err
	}

	if err := igw.Remove(c, vpcIds); err != nil {
		return err
	}

	if err := subnet.Remove(c, vpcIds); err != nil {
		return err
	}

	if err := rtb.Remove(c, vpcIds); err != nil {
		return err
	}

	if err := acl.Remove(c, vpcIds); err != nil {
		return err
	}

	if err := sg.Remove(c, vpcIds); err != nil {
		return err
	}

	if err := vpc.Remove(c, vpcIds); err != nil {
		return err
	}

	return nil
}
