package main

import (
	envcmd "FramerCLI/envCmd"
	modulecmd "FramerCLI/moduleCmd"
	origincmd "FramerCLI/originCmd"
	rootcmd "FramerCLI/rootCmd"
	runnercmd "FramerCLI/runnerCmd"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// add cmds
	cmd := &cobra.Command{
		Use:   "framer",
		Short: "Framer CLI",
		Long:  "Framer Command Line Interface",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootcmd.AddTo(cmd)
	envcmd.AddTo(cmd)
	runnercmd.AddTo(cmd)
	origincmd.AddTo(cmd)
	modulecmd.AddTo(cmd)

	// execute
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
