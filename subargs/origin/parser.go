package origin

import "github.com/akamensky/argparse"

func AddParser(parser *argparse.Parser) {
	parser.NewCommand("origin", "make or manage origins")
}
