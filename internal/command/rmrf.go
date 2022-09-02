package command

import (
	"github.com/knakayama/dv/internal/executor"
	"github.com/spf13/cobra"
)

type rmrfCmd struct {
	cmd  *cobra.Command
	opts rmrfOpts
}

type rmrfOpts struct {
	yes bool
}

func newRmrfCmd() *rmrfCmd {
	root := &rmrfCmd{}
	cmd := &cobra.Command{
		Use:           "rmrf",
		Short:         "Remove default VPCs in all AWS regions",
		Long:          `This command removes default VPCs in all AWS regions.`,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return executor.RemoveVpcs(root.opts.yes)
		},
	}

	cmd.PersistentFlags().BoolVarP(&root.opts.yes, "yes", "y", false, "You agree with deletion without any prompt")
	root.cmd = cmd

	return root
}
