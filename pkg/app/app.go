package app

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/RodrigoScola/ktype/pkg/book"
	"github.com/RodrigoScola/ktype/pkg/display"
	filesessions "github.com/RodrigoScola/ktype/pkg/file_sessions"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

type Options struct {
	Book *book.Book
}

func sessionCommand(_ *Options) *cli.Command {
	return &cli.Command{Name: "session", Aliases: []string{"s"}, Usage: "Retrieves a session", Action: func(context *cli.Context) error {

		var sess_name string
		names, err := filesessions.GetSessionNames()

		if err != nil {
			return err
		}

		names = append(names, "create_session")

		form := display.NewSessionMenu(names, &sess_name)
		if err := form.Run(); err != nil {
			return err
		}

		sess, err := getSession(sess_name)

		if err != nil {
			return err
		}

        goodEntries := make([]book.Sentence, 0)

            //make this better
        for _,v := range *&sess.Book.Sentences {
            if len(v.Current.Letters) == 0 {
                goodEntries = append(goodEntries, v)
            }
        }

		mybook := book.NewBook(goodEntries)
        mybook.Session = sess
        fmt.Println(sess.Name)
        sess.Book = mybook
        fmt.Println(sess)

		m := display.New(mybook)
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

	}}
}

func getSession(sess_name string) (*filesessions.FileSession, error) {
	if strings.Compare(sess_name, "create_session") != 0 {
		sess, err := filesessions.GetSession(sess_name)
		if err != nil {
			return nil, err
		}
		return sess, nil

	}

	var sess_filepath string
	menu := display.NewCreateSessionMenu(&sess_filepath)
	if err := menu.Run(); err != nil {
		return nil, err
	}
	file, err := os.ReadFile(sess_filepath)
	if err != nil {
		return nil, err
	}
	sentences := strings.Split(string(file), "\n")
	bookSentences := []book.Sentence{}
	for i := range sentences {
		if len(sentences[i]) > 0 {
			bookSentences = append(bookSentences, book.NewUserSentence(sentences[i]))
		}
	}

	sess, err := filesessions.Save(filesessions.New(
		&sess_filepath, &sentences, &bookSentences,
	))
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func GetApp() (*cli.App, *Options, error) {
	opts := Options{}
	return &cli.App{
		Commands: []*cli.Command{fileCommand(&opts), sessionCommand(&opts)},
	}, &opts, nil
}
