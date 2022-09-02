package executor

import (
	"fmt"

	"github.com/knakayama/dv/internal/entity"
)

func remove(region string, yes bool) error {
	if !yes {
		// TODO: pretty print
		//nolint:forbidigo
		fmt.Printf("%s skipped...\n", region)
		return nil
	}

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

func RemoveVpc(region string, yes bool) error {
	if err := validateRegion(region); err != nil {
		return err
	}

	if err := remove(region, yes); err != nil {
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
