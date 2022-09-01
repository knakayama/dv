package service

import (
	"fmt"

	"github.com/knakayama/dv/internal/entity"
)

func ListRegions() error {
	output, err := entity.NewRegion(entity.NewClient()).List()
	if err != nil {
		return err
	}

	for _, region := range output.Regions {
		// TODO: To be pretty printed
		//nolint:forbidigo
		fmt.Println(*region.RegionName)
	}

	return nil
}
