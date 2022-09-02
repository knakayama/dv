package executor

import (
	"github.com/knakayama/dv/internal/entity"
	"github.com/knakayama/dv/internal/presenter"
)

func ListVpcs() error {
	regionVpc := make(map[string]string)

	out, err := entity.NewRegion(entity.NewDefaultClient()).List()
	if err != nil {
		return err
	}

	for _, region := range out.Regions {
		vpc, err := entity.NewVpc(entity.NewClient(*region.RegionName))
		if err != nil {
			return err
		}

		switch vpc.Id {
		case nil:
			regionVpc[*region.RegionName] = "NaN"
		default:
			regionVpc[*region.RegionName] = *vpc.Id
		}
	}

	presenter.TableFrom(regionVpc)

	return nil
}
