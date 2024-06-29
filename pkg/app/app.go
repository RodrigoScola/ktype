package app

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)


type Options  struct {
    Contents []string
}

func fileCommand(contents []string) *cli.Command {
    return &cli.Command{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "types a file",
				Action: func(cCtx *cli.Context) error {
					file, err := os.Open(cCtx.Args().First())
					if err != nil {
						log.Fatal(err)
					}
					defer file.Close()

					content, err := io.ReadAll(file)
					if err != nil {
						log.Fatal(err)
					}

					contents = strings.Split(string(content), "\n")
					return nil
				},
			}
}

func GetApp() (*cli.App, *Options, error) {
    var contents []string
	return &cli.App{
		Commands: []*cli.Command{ fileCommand(contents), },
	}, &Options{
        Contents: contents,
    }, nil
}
