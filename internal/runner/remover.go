package runner

import "github.com/knakayama/dv/internal/entity"

type Remover struct{}

func (r *Remover) Run(args []string) error {
	region := args[0]

	vpc, err := entity.NewVpc(entity.NewClient(region))
	if err != nil {
		return err
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
