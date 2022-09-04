package executor

import (
	"github.com/knakayama/dv/internal/entity"
)

func RemoveVpcs(yes bool) error {
	out, err := entity.NewRegion(entity.NewDefaultClient()).List()
	if err != nil {
		return err
	}

	for _, r := range out.Regions {
		vpc, err := entity.NewVpc(entity.NewClient(*r.RegionName))
		if err != nil {
			return err
		}

		if vpc.Id == nil {
			continue
		}

		if err := removeNetworkComponents(*r.RegionName, vpc, yes); err != nil {
			return err
		}
	}

	return nil
}
