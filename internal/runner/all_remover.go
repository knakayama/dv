package runner

import (
	"fmt"

	"github.com/knakayama/dv/internal/service"
)

type AllRemover struct{}

func (a *AllRemover) Run(_ []string) error {
	for _, client := range service.MakeClients() {
		for _, vpc := range service.ListDefaultVpcs(client) {
			//nolint:forbidigo
			fmt.Println(*vpc.VpcId)
			service.DeleteIgws(client, vpc)
			service.DeleteSubnets(client, vpc)
			service.DeleteRouteTables(client, vpc)
			service.DeleteAcls(client, vpc)
			service.DeleteSecurityGroups(client, vpc)
			service.DeleteVpc(client, vpc)
		}
	}

	// TODO: Return an error
	return nil
}
