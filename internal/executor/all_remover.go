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

	for _, r := range out.Regions {
		//nolint:forbidigo
		fmt.Printf("==> %s\n", *r.RegionName)
		vpc, err := entity.NewVpc(entity.NewClient(*r.RegionName))
		if err != nil {
			return err
		}

		if vpc.Id == nil {
			//nolint:forbidigo
			fmt.Println("no vpc, skipped...")
			continue
		}

		if err := removeNetworkComponents(vpc, yes); err != nil {
			return err
		}
	}

	return nil
}
