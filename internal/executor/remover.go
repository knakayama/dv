package executor

import (
	"github.com/knakayama/dv/internal/entity"
)

type Remover struct{}

func remove(region string) error {
	vpc, err := entity.NewVpc(entity.NewClient(region))
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

	return nil
}

func (r *Remover) Run(args []string) error {
	region := args[0]
	if err := validateRegion(region); err != nil {
		return err
	}

	if err := remove(region); err != nil {
		return err
	}

	return nil
}

func validateRegion(regionLike string) error {
	out, err := entity.NewRegion(entity.NewDefaultClient()).List()
	if err != nil {
		return err
	}

	for _, region := range out.Regions {
		if regionLike == *region.RegionName {
			return nil
		}
	}

	return errUnknownRegion
}
