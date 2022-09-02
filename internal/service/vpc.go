package service

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/knakayama/dv/internal/entity"
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
	vpc, err := entity.NewVpc(entity.NewDefaultClient())
	if err != nil {
		return err
	}

	if err := vpc.NewIgw().Remove(); err != nil {
		return err
	}

	if err := vpc.NewSubnet().Remove(); err != nil {
		return err
	}

	if err := vpc.NewRouteTable().Remove(); err != nil {
		return err
	}

	if err := vpc.NewAcl().Remove(); err != nil {
		return err
	}

	if err := vpc.NewSecurityGroup().Remove(); err != nil {
		return err
	}

	if err := vpc.Remove(); err != nil {
		return err
	}

	return nil
}

func ListVpcs() error {
	output, err := entity.NewRegion(entity.NewDefaultClient()).List()
	if err != nil {
		return err
	}

	for _, region := range output.Regions {
		vpc, err := entity.NewVpc(entity.NewClient(*region.RegionName))
		if err != nil {
			return err
		}
		// TODO: To be pretty printed
		//nolint:forbidigo
		fmt.Println(*region.RegionName)
		//nolint:forbidigo
		fmt.Println(*vpc.Id)
	}

	return nil
}
