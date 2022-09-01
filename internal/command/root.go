package command

import (
	"log"

	"github.com/spf13/cobra"
)

type rootCmd struct {
	cmd *cobra.Command
}

func Execute(args []string) {
	newRootCmd().Execute(args)
}

func (cmd *rootCmd) Execute(args []string) {
	cmd.cmd.SetArgs(args)

	if err := cmd.cmd.Execute(); err != nil {
		log.Fatalf("Failed to execute, %v", err)
	}
}

func newRootCmd() *rootCmd {
	root := &rootCmd{
		cmd: &cobra.Command{
			Use:   "dv",
			Short: "Remove AWS default VPC(s)",
			Long: `This command enables you to remove default VPC in all AWS regions.
Aside from that, you can remove a VPC in each region.
			`,
			Version:       "0.0.1",
			SilenceUsage:  true,
			SilenceErrors: true,
		},
	}

	root.cmd.AddCommand(
		newRmrfCmd().cmd,
	)

	return root
}
