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
			Use:           "rm",
			Short:         "Remove a default VPC in an AWS region",
			Long:          `This command removes a default VPC in an AWS region`,
			SilenceUsage:  true,
			SilenceErrors: true,
			Args:          cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return service.Execute()
			},
		},
	}

	return root
}
