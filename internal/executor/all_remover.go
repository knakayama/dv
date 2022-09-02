package executor

import "github.com/knakayama/dv/internal/entity"

type AllRemover struct{}

func (a *AllRemover) Run(_ []string) error {
	out, err := entity.NewRegion(entity.NewDefaultClient()).List()
	if err != nil {
		return err
	}

	for _, region := range out.Regions {
		if err := remove(*region.RegionName); err != nil {
			return err
		}
	}

	return nil
}
