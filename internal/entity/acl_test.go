package entity

import (
	"testing"

	"github.com/knakayama/dv/test"
	"github.com/stretchr/testify/assert"
)

func setupAclTest(_ *testing.T) func(_ *testing.T) {
	test.RemoveVpcs()
	test.CreateDefaultVpc()

	return func(_ *testing.T) {
		test.RemoveVpcs()
	}
}

func TestAclRemoveNoVpc(t *testing.T) {
	teardownTest := setupVpcTest(t)
	defer teardownTest(t)

	vpc, _ := NewVpc(NewDefaultClient())
	err := vpc.NewAcl().Remove()

	assert.NotNil(t, err)
}

// Default VPC always has network components such as ACL.
func TestAclRemoveAclsExist(t *testing.T) {
	teardownTest := setupAclTest(t)
	defer teardownTest(t)

	vpc, _ := NewVpc(NewDefaultClient())
	test.CreateAcls(vpc.Id)
	err := vpc.NewAcl().Remove()
	acls := test.ListAcls(vpc.Id)

	assert.Nil(t, err)
	assert.Empty(t, acls)
}
