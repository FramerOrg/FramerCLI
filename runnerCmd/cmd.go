package runnercmd

import "github.com/spf13/cobra"

func AddTo(root *cobra.Command) {
	runnerCmd := &cobra.Command{
		Use:   "runner",
		Short: "Runner management",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	runnerCmd.Flags().Bool("exit-on-finish", false, "Exit on Finish")
	runnerCmd.Flags().Bool("restart-on-error", false, "Restart on Error")
	runnerCmd.Flags().Int("restart-sleep", 0, "Restart Sleep Seconds")
	runnerCmd.Flags().Bool("restart-on-file-change", false, "Restart on File Change")
	runnerCmd.Flags().StringSlice("start", nil, "Start Runner")
	root.AddCommand(runnerCmd)
}
