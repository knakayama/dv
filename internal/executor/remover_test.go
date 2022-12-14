package executor

import (
	"os"
	"testing"

	"github.com/knakayama/dv/internal/entity"
	"github.com/knakayama/dv/test"
	"github.com/stretchr/testify/assert"
)

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

func TestRemoverRemoveVpcNoVpc(t *testing.T) {
	test.RemoveVpcs(entity.NewDefaultClient())
	err := RemoveVpc(test.ValidRegion(), true)

	assert.ErrorIs(t, err, entity.ErrVpcNotFound)
}

func TestRemoverRemoveVpcVpcExists(t *testing.T) {
	client := entity.NewDefaultClient()
	test.RemoveVpcs(client)
	test.CreateDefaultVpc(client)

	err := RemoveVpc(os.Getenv("DEFAULT_REGION"), true)

	assert.Nil(t, err)
}
