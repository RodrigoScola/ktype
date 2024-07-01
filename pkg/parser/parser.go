package parser

import (
	"os"

	"github.com/hellflame/argparse"
)

type Opts struct {
    Args []string
    File string
}

func GetOpts() (*Opts, error) {
	parser := argparse.NewParser("proj", "gets all the values", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	args := parser.Strings("args", "arguments", &argparse.Option{
		Positional: false,
		Required:   false,
		Default:    "",
	})

	file := parser.String("f", "file", &argparse.Option{
		Required: true,
		Default:  "",
	})


	err := parser.Parse(os.Args)

	if err != nil {
		return nil, err
	}
	return &Opts{
        File: *file,
        Args: *args,
	}, nil

}

