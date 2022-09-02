package entity

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type Acl struct {
	vpc *Vpc
}

func (a *Acl) ids() ([]*string, error) {
	var aclIds []*string

	out, err := a.vpc.Client.DescribeNetworkAcls(
		context.TODO(),
		&ec2.DescribeNetworkAclsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*a.vpc.Id},
				},
			},
		},
	)
	if err != nil {
		return []*string{}, err
	}

	for _, acl := range out.NetworkAcls {
		aclIds = append(aclIds, acl.NetworkAclId)
	}

	return aclIds, nil
}

func (a *Acl) Remove() error {
	aclIds, _ := a.ids()

	for _, aclId := range aclIds {
		//nolint:forbidigo
		fmt.Println(*aclId)
		_, err := a.vpc.Client.DeleteNetworkAcl(
			context.TODO(),
			&ec2.DeleteNetworkAclInput{
				NetworkAclId: aclId,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *Vpc) NewAcl() *Acl {
	return &Acl{
		vpc: v,
	}
}
