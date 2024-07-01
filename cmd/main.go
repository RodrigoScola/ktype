package main

import (
	"fmt"
	"io"
	"os"

	"github.com/RodrigoScola/ktype/pkg/book"
	"golang.org/x/term"
)
func main() {
	var pipedInput string = "asodif"
	// Check if there's piped input
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == 0 {
		// There's piped input, read it
		pipedInput, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading piped input:", err)
			return
		}
		fmt.Println("Piped input received:", string(pipedInput))
	}

	// After handling piped input, check if stdin is a terminal
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		fmt.Fprintln(os.Stderr, "Not running in a terminal.")
		return
	}

	// Now it's safe to switch to raw mode since we're in a terminal
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error switching to raw mode:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), state)

	mybook := book.NewBook()

	mybook.Add(book.NewUserSentence(string(pipedInput)))
	// mybook.Add(book.NewUserSentence("the second one"))

	// for {
	// 	sen := mybook.Current()

	// 	b := make([]byte, 1)
	// 	_, err = os.Stdin.Read(b)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	display.Clear()

	// 	if strings.Compare(string(b), "\x03") == 0 {
	// 		break
	// 	} else if strings.Compare(string(b), "\x7f") == 0 {
	// 		length := sen.Current.Length() - 1
	// 		sen.Current.Remove(length)
	// 	} else if strings.Compare(string(b), "\r") == 0 {
	// 		sen = mybook.Next()
	// 	} else {
	// 		sen.Current.Add(book.Letter{Char: b[0], Ignore: false})
	// 	}
	// 	if sen.Complete() {
	// 		sen = mybook.Next()
	// 	}

	// 	fmt.Printf("%q", b)

	// 	fmt.Println(sen.Display())
	// }
}
