package cmd

import (
	"fmt"

	"github.com/knakayama/dv/internal"
	"github.com/spf13/cobra"
)

type rmAllCmd struct {
	cmd *cobra.Command
}

func newRmAllCmd() *rmAllCmd {
	root := &rmAllCmd{
		cmd: &cobra.Command{
			Use:           "rm-all",
			Short:         "Remove default VPCs in all AWS regions",
			Long:          `Remove default VPCs in all AWS regions`,
			Aliases:       []string{"a"},
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
