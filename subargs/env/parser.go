package env

import "github.com/akamensky/argparse"

func AddParser(parser *argparse.Parser) {
	parser.NewCommand("env", "manage env.json file")
}
