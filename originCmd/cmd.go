package origincmd

import "github.com/spf13/cobra"

func AddTo(root *cobra.Command) {
	originCmd := &cobra.Command{
		Use:   "origin",
		Short: "Origin management",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	originCmd.Flags().String("add", "", "Add Origin")
	originCmd.Flags().BoolP("list", "l", false, "List Origins")
	originCmd.Flags().String("del", "", "Delete Origin")
	originCmd.Flags().Bool("sync", false, "Sync Origin")
	originCmd.Flags().Bool("make", false, "Make Origin")
	root.AddCommand(originCmd)
}
