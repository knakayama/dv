package entity

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type Vpc struct {
	// You cannot have more than one default VPC per Region.
	// https://docs.aws.amazon.com/cli/latest/reference/ec2/create-default-vpc.html
	Id     *string
	Client *ec2.Client
}

func (v *Vpc) Remove() error {
	if v.Id == nil {
		return ErrVpcNotFound
	}

	_, err := v.Client.DeleteVpc(
		context.TODO(),
		&ec2.DeleteVpcInput{
			VpcId: v.Id,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func NewVpc(client *ec2.Client) (*Vpc, error) {
	out, err := client.DescribeVpcs(
		context.TODO(),
		&ec2.DescribeVpcsInput{})
	if err != nil {
		return &Vpc{}, err
	}

	for _, vpc := range out.Vpcs {
		if *vpc.IsDefault {
			return &Vpc{Id: vpc.VpcId, Client: client}, nil
		}
	}

	return &Vpc{}, nil
}
