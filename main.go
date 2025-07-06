package main

import (
	"fmt"
	"os"

	"FramerCLI/subargs/env"
	"FramerCLI/subargs/module"
	"FramerCLI/subargs/origin"
	"FramerCLI/subargs/root"
	"FramerCLI/subargs/runner"

	"github.com/akamensky/argparse"
)

func main() {
	// create parsers
	rootParser := argparse.NewParser("FramerCLI", "Framer Official CLI")
	envParser := rootParser.NewCommand("env", "manage env.json file")
	runnerParser := rootParser.NewCommand("runner", "cli build-in python runner and watcher")
	originParser := rootParser.NewCommand("origin", "make or manage origins")
	moduleParser := rootParser.NewCommand("module", "create or manage modules")

	// add sub args
	root.AddArgs(rootParser)
	env.AddArgs(envParser)
	runner.AddArgs(runnerParser)
	origin.AddArgs(originParser)
	module.AddArgs(moduleParser)

	// parse arguments
	err := rootParser.Parse(os.Args)
	if err != nil {
		fmt.Println(rootParser.Usage(err))
		return
	}
}
