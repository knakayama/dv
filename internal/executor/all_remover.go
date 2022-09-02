package executor

import "github.com/knakayama/dv/internal/entity"

type AllRemover struct{}

func (a *AllRemover) Run(_ []string) error {
	out, err := entity.NewRegion(entity.NewDefaultClient()).List()
	if err != nil {
		return err
	}

	for _, region := range out.Regions {
		vpc, err := entity.NewVpc(entity.NewClient(*region.RegionName))
		if err != nil {
			return err
		}

		if vpc.Id == nil {
			return errVpcNotFound
		}

		if err := vpc.NewIgw().Remove(); err != nil {
			return err
		}

		if err := vpc.NewSubnet().Remove(); err != nil {
			return err
		}

		if err := vpc.NewRouteTable().Remove(); err != nil {
			return err
		}

		if err := vpc.NewAcl().Remove(); err != nil {
			return err
		}

		if err := vpc.NewSecurityGroup().Remove(); err != nil {
			return err
		}

		if err := vpc.Remove(); err != nil {
			return err
		}
	}
	return nil
}
