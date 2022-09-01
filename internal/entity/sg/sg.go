package sg

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/knakayama/dv/internal/util"
)

func ids(client *ec2.Client, vpcIds []*string) ([]*string, error) {
	var sgIds []*string

	output, err := client.DescribeSecurityGroups(
		context.TODO(),
		&ec2.DescribeSecurityGroupsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: util.StrListFrom(vpcIds),
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

func Remove(client *ec2.Client, vpcIds []*string) error {
	sgIds, _ := ids(client, vpcIds)

	for _, sgId := range sgIds {
		//nolint:forbidigo
		fmt.Println(*sgId)
		_, err := client.DeleteSecurityGroup(
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
