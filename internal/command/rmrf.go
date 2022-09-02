package command

import (
	"github.com/knakayama/dv/internal/runner"
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
			Long:          `This command removes default VPCs in all AWS regions.`,
			SilenceUsage:  true,
			SilenceErrors: true,
			Args:          cobra.NoArgs,
			RunE:          runE(&runner.AllRemover{}),
		},
	}

	return root
}
