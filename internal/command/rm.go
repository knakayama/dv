package command

import (
	"github.com/knakayama/dv/internal/executor"
	"github.com/spf13/cobra"
)

type rmCmd struct {
	cmd  *cobra.Command
	opts rmCmdOpts
}

type rmCmdOpts struct {
	yes bool
}

func newRmCmd() *rmCmd {
	root := &rmCmd{}
	cmd := &cobra.Command{
		Use:           "rm [region]",
		Short:         "Remove a default VPC in an AWS region",
		Long:          `This command removes a default VPC in an AWS region`,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return executor.RemoveVpc(args[0], root.opts.yes)
		},
	}

	cmd.PersistentFlags().BoolVarP(&root.opts.yes, "yes", "y", false, "You agree with deletion without any prompt")
	root.cmd = cmd

	return root
}
