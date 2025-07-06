package envcmd

import "github.com/spf13/cobra"

func AddTo(root *cobra.Command) {
	envCmd := &cobra.Command{
		Use:   "env",
		Short: "Environment management",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	envCmd.Flags().Bool("init", false, "Init Env File")
	envCmd.Flags().BoolP("list", "l", false, "List Environments")
	envCmd.Flags().StringSlice("set", nil, "Set Environment, TYPE can be 'str', 'int', 'float', 'bool', Default 'str'")
	envCmd.Flags().String("del", "", "Delete Environment")
	root.AddCommand(envCmd)
}
