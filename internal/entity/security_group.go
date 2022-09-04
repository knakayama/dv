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

func (s *SecurityGroup) ids() ([]*string, error) {
	var sgIds []*string

	out, err := s.vpc.Client.DescribeSecurityGroups(
		context.TODO(),
		&ec2.DescribeSecurityGroupsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*s.vpc.Id},
				},
			},
		},
	)
	if err != nil {
		return []*string{}, err
	}

	for _, sg := range out.SecurityGroups {
		sgIds = append(sgIds, sg.GroupId)
	}

	return sgIds, nil
}

func (s *SecurityGroup) Remove() error {
	if s.vpc.Id == nil {
		return nil
	}

	sgIds, _ := s.ids()

	for _, sgId := range sgIds {
		//nolint:forbidigo
		fmt.Println(*sgId)
		_, err := s.vpc.Client.DeleteSecurityGroup(
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

func (v *Vpc) NewSecurityGroup() *SecurityGroup {
	return &SecurityGroup{
		vpc: v,
	}
}
