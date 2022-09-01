package service

import (
	"github.com/knakayama/dv/internal/entity"
)

func Execute() error {
	client := entity.NewClient()

	vpc, err := entity.NewVpc(client)
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
