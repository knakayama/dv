package cmd

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
			Use:           "rmadv",
			Short:         "Remove AWS default VPCs",
			Long:          `Remove AWS default VPCs`,
			Version:       "0.0.1",
			SilenceUsage:  true,
			SilenceErrors: true,
		},
	}

	root.cmd.AddCommand(
		newRmAllCmd().cmd,
	)

	return root
}
