package app

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/RodrigoScola/ktype/pkg/book"
	"github.com/RodrigoScola/ktype/pkg/display"
	"github.com/RodrigoScola/ktype/pkg/sessions"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

type Options struct {
	Book *book.Book
}

func sessionCommand(_ *Options) *cli.Command {
	return &cli.Command{Name: "session", Aliases: []string{"s"}, Usage: "Retrieves a session", Action: func(context *cli.Context) error {
		fmt.Println("")
		sessions.Setup()
		all_names, err := os.ReadDir(filepath.Join(".", "data", "sessions"))
		if err != nil {
			return err
		}
		names := make([]string, 0)
		for _, entry := range all_names {
			if entry.IsDir() {
				continue
			}
			names = append(names, entry.Name())
		}

		names = append(names, "create_session")

		var sessionname string
		p := display.NewSessionMenu(names, &sessionname)

		if err := p.Run(); err != nil {
			log.Fatalf("err : %w", err)
		}
		var sess *sessions.Session
		if strings.Compare(sessionname, "create_session") == 0 {
			var new_session_name string

			s := display.NewCreateSessionMenu(&new_session_name)

			if err := s.Run(); err != nil {
				log.Fatalf("err : %w", err)
			}
			fmt.Println(new_session_name)
			ex, err := sessions.Exists(new_session_name)
			if err != nil {
				return err
			}
			if *ex == false {
				createNewSession(new_session_name)
			} else {
				sess, err = sessions.GetSession(new_session_name)
				if err != nil {
					return err
				}
				fmt.Println(sess)
			}
		} else {
			sess, err = sessions.GetSession(sessionname)
			if err != nil {
				return err
			}
		}
		mybook := book.NewBook()
		for _, str := range sess.File {
			if len(str) == 0 {
				continue
			}
			mybook.Add(book.NewUserSentence(str))
		}


		m := display.New(sessions.NewTypingSession(mybook))
		m.Input.Focus()
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		mn := tea.NewProgram(m, tea.WithAltScreen())

		if _, err := mn.Run(); err != nil {
			log.Fatalf("err : %w", err)
		}

		return nil
	}}
}

func createNewSession(session_name string) error {
	file, err := os.Open(session_name)
	if err != nil {
		return err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	sess := sessions.NewSession(session_name, strings.Split(string(content), "\n"))
	fmt.Println(sess)
	sessions.Save(sess)
	return nil
}

func GetApp() (*cli.App, *Options, error) {
	opts := Options{}
	return &cli.App{
		Commands: []*cli.Command{fileCommand(&opts), sessionCommand(&opts)},
	}, &opts, nil
}
