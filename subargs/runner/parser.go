package runner

import "github.com/akamensky/argparse"

func AddParser(parser *argparse.Parser) {
	parser.NewCommand("runner", "cli build-in python runner and watcher")
}
