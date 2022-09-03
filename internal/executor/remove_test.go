package executor

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/knakayama/dv/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestRemoverValidateRegionInvalidRegion(t *testing.T) {
	faker := faker.New()
	err := validateRegion(faker.RandomLetter())

	assert.ErrorIs(t, err, entity.ErrUnknownRegion)
}

func TestRemoverValidateRegionValidRegion(t *testing.T) {
	for _, tt := range []string{
		"af-south-1",
		"eu-north-1",
		"ap-south-1",
		"eu-west-3",
		"eu-west-2",
		"eu-south-1",
		"eu-west-1",
		"ap-northeast-3",
		"ap-northeast-2",
		"me-south-1",
		"ap-northeast-1",
		"me-central-1",
		"sa-east-1",
		"ca-central-1",
		"ap-east-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-southeast-3",
		"eu-central-1",
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
	} {
		err := validateRegion(tt)

		assert.Nil(t, err)
	}
}
