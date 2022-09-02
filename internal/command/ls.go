package command

import (
	"github.com/knakayama/dv/internal/executor"
	"github.com/spf13/cobra"
)

type lsCmd struct {
	cmd *cobra.Command
}

func newlsCmd() *lsCmd {
	root := &lsCmd{
		cmd: &cobra.Command{
			Use:           "ls",
			Short:         "List VPCs in all AWS regions",
			Long:          `This command lists VPCs in all AWS regions`,
			SilenceUsage:  true,
			SilenceErrors: true,
			Args:          cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return executor.ListVpcs()
			},
		},
	}

	return root
}
