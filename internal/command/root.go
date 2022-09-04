package command

import (
	"log"

	"github.com/spf13/cobra"
)

type rootCmd struct {
	cmd  *cobra.Command
	exit func(int)
}

func Execute(args []string, exit func(int)) {
	newRootCmd(exit).Execute(args)
}

func (cmd *rootCmd) Execute(args []string) {
	cmd.cmd.SetArgs(args)

	if err := cmd.cmd.Execute(); err != nil {
		log.Fatalf("Failed to execute, %v", err)
		// TODO: Add a feature to pass an appropriate code from callees
		cmd.exit(1)
	}
}

func newRootCmd(exit func(int)) *rootCmd {
	root := &rootCmd{
		cmd: &cobra.Command{
			Use:   "dv",
			Short: "Remove AWS default VPC(s)",
			Long: `This command enables you to remove default VPC(s) in all AWS regions.
Aside from that, you can remove a VPC in each region.
			`,
			Version:       "0.0.1",
			SilenceUsage:  true,
			SilenceErrors: true,
			Args:          cobra.NoArgs,
		},
		exit: exit,
	}

	root.cmd.AddCommand(
		newRmrfCmd().cmd,
		newRmCmd().cmd,
		newLsCmd().cmd,
	)

	return root
}
