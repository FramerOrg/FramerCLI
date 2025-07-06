package modulecmd

import "github.com/spf13/cobra"

func AddTo(root *cobra.Command) {
	moduleCmd := &cobra.Command{
		Use:   "module",
		Short: "Module management",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	moduleCmd.Flags().BoolP("list", "l", false, "List Modules")
	moduleCmd.Flags().Bool("sync-pkg", false, "Sync Installed Package To Framerpkg")
	moduleCmd.Flags().String("info", "", "Show Module Info")
	moduleCmd.Flags().String("enable", "", "Enable Module")
	moduleCmd.Flags().String("disable", "", "Disable Module")
	moduleCmd.Flags().String("del", "", "Delete Module")
	moduleCmd.Flags().StringP("search", "s", "", "Search Module")
	moduleCmd.Flags().Bool("overwrite", false, "Install Module and Override Old")
	moduleCmd.Flags().StringP("install", "i", "", "Install Module")
	moduleCmd.Flags().Bool("sync-back", false, "Sync Package Back")
	moduleCmd.Flags().String("create", "", "Create Empty Framer Modules")
	root.AddCommand(moduleCmd)
}
