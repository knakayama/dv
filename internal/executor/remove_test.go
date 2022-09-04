package executor

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/knakayama/dv/internal/entity"
	"github.com/knakayama/dv/test"
	"github.com/stretchr/testify/assert"
)

func setup(_ *testing.T, client *ec2.Client) func(_ *testing.T) {
	test.RemoveVpcs(client)
	test.CreateDefaultVpc(client)

	return func(_ *testing.T) {
		test.RemoveVpcs(client)
	}
}

func TestRemoverValidateRegionInvalidRegion(t *testing.T) {
	err := validateRegion(test.InvalidRegion())

	assert.ErrorIs(t, err, entity.ErrUnknownRegion)
}

func TestRemoverValidateRegionValidRegion(t *testing.T) {
	for _, tt := range test.ValidRegions() {
		err := validateRegion(tt)

		assert.Nil(t, err)
	}
}

func TestRemoveRemoverNetworkComponentsYesIsFalse(t *testing.T) {
	err := removeNetworkComponents(test.InvalidRegion(), &entity.Vpc{}, false)

	assert.Nil(t, err)
}

func TestRemoveRemoveVpcNoVpc(t *testing.T) {
	test.RemoveVpcs(entity.NewDefaultClient())
	err := RemoveVpc(test.ValidRegion(), true)

	assert.ErrorIs(t, err, entity.ErrVpcNotFound)
}

func TestRemoveRemoveVpcVpcExists(t *testing.T) {
	client := entity.NewDefaultClient()
	teardown := setup(t, client)
	defer teardown(t)

	err := RemoveVpc(test.ValidRegion(), true)

	assert.Nil(t, err)
}
