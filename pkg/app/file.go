package app

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/RodrigoScola/ktype/pkg/book"
	"github.com/RodrigoScola/ktype/pkg/display"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)


func fileCommand(opts *Options) *cli.Command {
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

			opts.Book = book.NewBook([]book.Sentence{})
			for _, str := range strings.Split(string(content), "\n") {
				if len(str) == 0 {
					continue
				}
				opts.Book.Add(book.NewUserSentence(str))
			}

			m := display.New(opts.Book)
			m.Input.Focus()
			f, err := tea.LogToFile("debug.log", "debug")
			if err != nil {
				panic(err)
			}
			defer f.Close()

			p := tea.NewProgram(m, tea.WithAltScreen())

			if _, err := p.Run(); err != nil {
				log.Fatalf("err : %w", err)
			}
			return nil
		},
	}
}
