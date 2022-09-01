package command

import (
	"fmt"

	"github.com/knakayama/dv/internal/service"
	"github.com/spf13/cobra"
)

type rmrfCmd struct {
	cmd *cobra.Command
}

func newRmrfCmd() *rmrfCmd {
	root := &rmrfCmd{
		cmd: &cobra.Command{
			Use:           "rmrf",
			Short:         "Remove default VPCs in all AWS regions",
			Long:          `This command remove default VPCs in all AWS regions.`,
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
