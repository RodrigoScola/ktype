package app

import (
	"fmt"

	"github.com/RodrigoScola/ktype/pkg/book"
	filesessions "github.com/RodrigoScola/ktype/pkg/file_sessions"
	"github.com/RodrigoScola/ktype/pkg/statistics"
	"github.com/urfave/cli/v2"
)

type Options struct {
	Book *book.Book
}

func profileCommand(_ *Options) *cli.Command {
	return &cli.Command{Name: "profile", Aliases: []string{"p"}, Usage: "Shows the profile info", Action: func(context *cli.Context) error {
		fmt.Println("The Profile Info")

		if err := calc(); err != nil {
			return err
		}

		return nil

	}}
}

func calc() error {
	names, err := filesessions.GetSessionNames()

	if err != nil {
		return err
	}
	for _, name := range names {
		sess, err := filesessions.GetSession(name)
		if err != nil {
			return err
		}
		for _, v := range sess.Sentences {
			if v.Current.GetFirst() == nil && v.Current.GetLast() == nil {
				continue
			}
			fmt.Printf("%.2f\n", statistics.CalculateWPM(v.Current.GetFirst().CreatedAt,
				v.Current.GetLast().CreatedAt, len(v.Current.Letters)))

		}

	}

	return nil
}

func GetApp() (*cli.App, *Options, error) {
	opts := Options{}
	return &cli.App{
		Commands: []*cli.Command{fileCommand(&opts), sessionCommand(&opts), profileCommand(&opts)},
	}, &opts, nil
}
