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

func deleteIgws(client *ec2.Client, vpc types.Vpc) {
	output, err := client.DescribeInternetGateways(
		context.TODO(),
		&ec2.DescribeInternetGatewaysInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("attachment.vpc-id"),
					Values: []string{*vpc.VpcId},
				},
			},
		},
	)

	if err != nil {
		log.Fatalf("Failed to list igws, %v", err)
	}

	for _, igw := range output.InternetGateways {
		fmt.Println(*igw.InternetGatewayId)
		_, err := client.DeleteInternetGateway(
			context.TODO(),
			&ec2.DeleteInternetGatewayInput{
				InternetGatewayId: igw.InternetGatewayId,
				DryRun:            aws.Bool(true),
			},
		)

		if err != nil {
			log.Fatalf("Failed to delete an igw, %v", err)
		}
	}
}

func deleteSubnets(client *ec2.Client, vpc types.Vpc) {
	output, err := client.DescribeSubnets(
		context.TODO(),
		&ec2.DescribeSubnetsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*vpc.VpcId},
				},
			},
		},
	)

	if err != nil {
		log.Fatalf("Failed to list subnets, %v", err)
	}

	for _, subnet := range output.Subnets {
		fmt.Println(*subnet.SubnetId)
		_, err := client.DeleteSubnet(
			context.TODO(),
			&ec2.DeleteSubnetInput{
				SubnetId: subnet.SubnetId,
				DryRun:   aws.Bool(true),
			},
		)

		if err != nil {
			log.Fatalf("Failed to delete a subnet, %v", err)
		}
	}
}

func deleteRouteTables(client *ec2.Client, vpc types.Vpc) {
	output, err := client.DescribeRouteTables(
		context.TODO(),
		&ec2.DescribeRouteTablesInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*vpc.VpcId},
				},
			},
		},
	)

	if err != nil {
		log.Fatalf("Failed to list route tables, %v", err)
	}

	for _, routeTable := range output.RouteTables {
		fmt.Println(*routeTable.RouteTableId)
		_, err := client.DeleteRouteTable(
			context.TODO(),
			&ec2.DeleteRouteTableInput{
				RouteTableId: routeTable.RouteTableId,
				DryRun:       aws.Bool(true),
			},
		)

		if err != nil {
			log.Fatalf("Failed to delete a route table, %v", err)
		}
	}
}

func deleteAcls(client *ec2.Client, vpc types.Vpc) {
	output, err := client.DescribeNetworkAcls(
		context.TODO(),
		&ec2.DescribeNetworkAclsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*vpc.VpcId},
				},
			},
		},
	)

	if err != nil {
		log.Fatalf("Failed to list network acls, %v", err)
	}

	for _, acl := range output.NetworkAcls {
		fmt.Println(*acl.NetworkAclId)
		_, err := client.DeleteNetworkAcl(
			context.TODO(),
			&ec2.DeleteNetworkAclInput{
				NetworkAclId: acl.NetworkAclId,
				DryRun:       aws.Bool(true),
			},
		)

		if err != nil {
			log.Fatalf("Failed to delete a network acl, %v", err)
		}
	}
}

func deleteSecurityGroups(client *ec2.Client, vpc types.Vpc) {
	output, err := client.DescribeSecurityGroups(
		context.TODO(),
		&ec2.DescribeSecurityGroupsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*vpc.VpcId},
				},
			},
		},
	)

	if err != nil {
		log.Fatalf("Failed to list security groups, %v", err)
	}

	for _, securityGroup := range output.SecurityGroups {
		_, err := client.DeleteSecurityGroup(
			context.TODO(),
			&ec2.DeleteSecurityGroupInput{
				DryRun:  aws.Bool(true),
				GroupId: securityGroup.GroupId,
			},
		)

		if err != nil {
			log.Fatalf("Failed to a security groups, %v", err)
		}
	}
}

func deleteVpc(client *ec2.Client, vpc types.Vpc) {
	_, err := client.DeleteVpc(
		context.TODO(),
		&ec2.DeleteVpcInput{
			VpcId:  vpc.VpcId,
			DryRun: aws.Bool(true),
		},
	)

	if err != nil {
		log.Fatalf("Failed to delete a vpc, %v", err)
	}
}

func main() {
	for _, client := range makeClients() {
		for _, vpc := range listDefaultVpcs(client) {
			fmt.Println(*vpc.VpcId)
			deleteIgws(client, vpc)
			deleteSubnets(client, vpc)
			deleteRouteTables(client, vpc)
			deleteAcls(client, vpc)
			deleteSecurityGroups(client, vpc)
			deleteVpc(client, vpc)
		}
	}
}
