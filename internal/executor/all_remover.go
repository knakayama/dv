package executor

import (
	"github.com/knakayama/dv/internal/entity"
)

func RemoveVpcs(yes bool) error {
	out, err := entity.NewRegion(entity.NewDefaultClient()).List()
	if err != nil {
		return err
	}

	for _, region := range out.Regions {
		if err := remove(*region.RegionName, yes); err != nil {
			return err
		}
	}

	return nil
}
