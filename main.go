package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/RodrigoScola/ktype/pkg/book"
	"github.com/RodrigoScola/ktype/pkg/display"
	cli "github.com/urfave/cli/v2"
	"golang.org/x/term"
)

func main() {
	//	display.InitTea()

    var contents []string

	app := &cli.App{
		Commands: []*cli.Command{
			{
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

                    contents= strings.Split(string(content), "\n")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}



	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), state)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	   mybook := book.NewBook()


    for _, value:= range contents{
	   mybook.Add(book.NewUserSentence(value))

    }


	fmt.Println(mybook.Current().Display())

	   for {
	       sen := mybook.Current()

	       b := make([]byte, 1)
	       _, err = os.Stdin.Read(b)
	       if err != nil {
	           panic(err)
	       }
	       display.Clear()

	       if strings.Compare(string(b), "\x03") == 0 {
	           break
	       } else if strings.Compare(string(b), "\x7f") == 0 {
            sen.Current.RemoveLast()
	       } else if strings.Compare(string(b), "\r") == 0 {
	           sen = mybook.Next();
	       } else {
	               sen.Current.Add(book.Letter{Char: b[0], Ignore: false})
	       }
	       if sen.Complete()  {
	           sen = mybook.Next();
	       }

	       //fmt.Printf("%q", b)


	       fmt.Println(sen.Display())
	   }
}
