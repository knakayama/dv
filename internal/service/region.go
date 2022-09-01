package service

import (
	"fmt"

	"github.com/knakayama/dv/internal/entity/client"
	"github.com/knakayama/dv/internal/entity/region"
)

func ListRegions() error {
	output, err := region.List((client.New()))
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
