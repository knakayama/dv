package command

import (
	"fmt"

	"github.com/knakayama/dv/internal/service"
	"github.com/spf13/cobra"
)

type rmCmd struct {
	cmd *cobra.Command
}

func newRmCmd() *rmCmd {
	root := &rmCmd{
		cmd: &cobra.Command{
			Use:           "rm",
			Short:         "Remove a default VPC in an AWS region",
			Long:          `This command removes a default VPC in an AWS region`,
			SilenceUsage:  true,
			SilenceErrors: true,
			Args:          cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
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
			},
		},
	}

	return root
}
