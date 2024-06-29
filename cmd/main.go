package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/RodrigoScola/ktype/pkg/book"
	"github.com/RodrigoScola/ktype/pkg/display"
	"golang.org/x/term"
)

func main() {
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), state)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

    mybook := book.NewBook()


    mybook.Add(book.NewUserSentence("this should be the correct thing"))
    mybook.Add(book.NewUserSentence("the second one"))

	// sen.Current.Add(sentence.Letter{
	// 	Char:   't',
	// 	Ignore: false,
	// })
	// sen.Current.Add(sentence.Letter{
	// 	Char:   'h',
	// 	Ignore: false,
	// })
	// sen.Current.Add(sentence.Letter{
	// 	Char:   'z',
	// 	Ignore: false,
	// })
	// sen.Current.Add(sentence.Letter{
	// 	Char:   's',
	// 	Ignore: false,
	// })

	//fmt.Println(sen.Compare())
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
            length := sen.Current.Length() -1
            sen.Current.Remove(length)
        } else if strings.Compare(string(b), "\r") == 0 {
            sen = mybook.Next();
        } else {
                sen.Current.Add(book.Letter{Char: b[0], Ignore: false})
        }
        if sen.Complete()  {
            sen = mybook.Next();
        }

        fmt.Printf("%q", b)


        fmt.Println(sen.Display())
    }
}
