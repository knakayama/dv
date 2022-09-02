package executor

import (
	"fmt"

	"github.com/knakayama/dv/internal/entity"
)

func RemoveVpcs(yes bool) error {
	out, err := entity.NewRegion(entity.NewDefaultClient()).List()
	if err != nil {
		return err
	}

	for _, region := range out.Regions {
		if yes {
			if err := remove(*region.RegionName); err != nil {
				return err
			}
		}
		// TODO: pretty print
		//nolint:forbidigo
		fmt.Printf("%s skipped...\n", *region.RegionName)
		continue
	}

	return nil
}
