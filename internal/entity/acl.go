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

func (acl *Acl) ids() ([]*string, error) {
	var aclIds []*string

	out, err := acl.vpc.Client.DescribeNetworkAcls(
		context.TODO(),
		&ec2.DescribeNetworkAclsInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("vpc-id"),
					Values: []string{*acl.vpc.Id},
				},
			},
		},
	)
	if err != nil {
		return []*string{}, err
	}

	for _, acl := range out.NetworkAcls {
		// We can't delete Default ACL
		if *acl.IsDefault {
			//nolint:forbidigo
			fmt.Printf("%s is Default ACL, skipped...\n", *acl.NetworkAclId)
			continue
		}
		aclIds = append(aclIds, acl.NetworkAclId)
	}

	return aclIds, nil
}

func (acl *Acl) Remove() error {
	if acl.vpc.Id == nil {
		return nil
	}

	aclIds, _ := acl.ids()

	for _, aclId := range aclIds {
		//nolint:forbidigo
		fmt.Println(*aclId)
		_, err := acl.vpc.Client.DeleteNetworkAcl(
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
