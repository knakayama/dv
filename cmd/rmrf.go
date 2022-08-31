package cmd

import (
	"fmt"

	"github.com/knakayama/dv/internal"
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
				for _, client := range internal.MakeClients() {
					for _, vpc := range internal.ListDefaultVpcs(client) {
						//nolint:forbidigo
						fmt.Println(*vpc.VpcId)
						internal.DeleteIgws(client, vpc)
						internal.DeleteSubnets(client, vpc)
						internal.DeleteRouteTables(client, vpc)
						internal.DeleteAcls(client, vpc)
						internal.DeleteSecurityGroups(client, vpc)
						internal.DeleteVpc(client, vpc)
					}
				}
			},
		},
	}

	return root
}
