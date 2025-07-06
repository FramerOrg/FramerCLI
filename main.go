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
	parser := argparse.NewParser("FramerCLI", "Framer Official CLI")
	root.AddParser(parser)
	env.AddParser(parser)
	runner.AddParser(parser)
	origin.AddParser(parser)
	module.AddParser(parser)

	// parse arguments
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}
}
