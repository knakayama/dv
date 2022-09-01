package entity

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type SecurityGroup struct {
	vpc *Vpc
}

func (sg *SecurityGroup) ids() ([]*string, error) {
	var sgIds []*string

	output, err := sg.vpc.Client.DescribeSecurityGroups(
		context.TODO(),
		&ec2.DescribeSecurityGroupsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*sg.vpc.Id},
				},
			},
		},
	)
	if err != nil {
		return []*string{}, err
	}

	for _, sg := range output.SecurityGroups {
		sgIds = append(sgIds, sg.GroupId)
	}

	return sgIds, nil
}

func (sg *SecurityGroup) Remove() error {
	sgIds, _ := sg.ids()

	for _, sgId := range sgIds {
		//nolint:forbidigo
		fmt.Println(*sgId)
		_, err := sg.vpc.Client.DeleteSecurityGroup(
			context.TODO(),
			&ec2.DeleteSecurityGroupInput{
				GroupId: sgId,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (vpc *Vpc) NewSecurityGroup() *SecurityGroup {
	return &SecurityGroup{
		vpc: vpc,
	}
}
