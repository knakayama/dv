package command

import (
	"github.com/knakayama/dv/internal/service"
	"github.com/spf13/cobra"
)

type rmCmd struct {
	cmd *cobra.Command
}

func newRmCmd() *rmCmd {
	root := &rmCmd{
		cmd: &cobra.Command{
			Use:           "rm [region]",
			Short:         "Remove a default VPC in an AWS region",
			Long:          `This command removes a default VPC in an AWS region`,
			SilenceUsage:  true,
			SilenceErrors: true,
			Args:          cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				return service.RemoveVpc(args[0])
			},
		},
	}

	return root
}
