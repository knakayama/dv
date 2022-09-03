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

	for _, r := range out.Regions {
		vpc, err := entity.NewVpc(entity.NewClient(*r.RegionName))
		if err != nil {
			return err
		}

		switch vpc.Id {
		case nil:
			regionVpc[*r.RegionName] = "NaN"
		default:
			regionVpc[*r.RegionName] = *vpc.Id
		}
	}

	presenter.TableFrom(regionVpc, [2]string{"Region", "Default VPC"})

	return nil
}
