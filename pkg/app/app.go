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

        fmt.Println("sessions -> ")
        if err := calc(); err != nil {
            return err
        }

		return nil

	}}
}

func calc() error {
    names, err := filesessions.GetSessionNames()
    if err != nil { return err }
    books := []*book.Book{}
    wpms := []float64{}
    for i := range names {
        b , err := filesessions.GetSession(names[i])
        if err != nil { return err }
        books = append(books, b)
        filled := b.FilledSentences()

        for k := range filled {
            wpm := statistics.CalculateWPM(
                filled[k].Current.Letters[0].CreatedAt,
                filled[k].Current.GetLast().CreatedAt,
                    len(filled[k].Correct),
                )
            wpms = append(wpms, wpm)
        }
    }
    fmt.Println(wpms)
    return nil
}



func GetApp() (*cli.App, *Options, error) {
	opts := Options{}
	return &cli.App{
		Commands: []*cli.Command{fileCommand(&opts), sessionCommand(&opts), profileCommand(&opts)},
	}, &opts, nil
}
