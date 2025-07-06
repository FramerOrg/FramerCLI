package rootcmd

import "github.com/spf13/cobra"

func AddTo(cmd *cobra.Command) {
	cmd.Flags().BoolP("version", "v", false, "Show Version")
	cmd.Flags().Bool("shell", false, "Open Framer Shell")
	cmd.Flags().BoolP("test", "t", false, "Test Framer")
	cmd.Flags().Bool("init", false, "Init Project")
	cmd.Flags().StringSliceP("module", "m", nil, "Load Module CLI")
	cmd.Flags().Bool("update", false, "Update Framer")
}
