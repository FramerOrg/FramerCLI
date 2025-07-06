package module

import "github.com/akamensky/argparse"

func AddParser(parser *argparse.Parser) {
	parser.NewCommand("module", "create or manage modules")
}
