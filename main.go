package main

import (
	"log"
	"os"


	"github.com/RodrigoScola/ktype/pkg/app"
)

func main() {
	app, _, err := app.GetApp()
	if err != nil {
		panic(err)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	// state, err := term.MakeRaw(int(os.Stdin.Fd()))
	// defer term.Restore(int(os.Stdin.Fd()), state)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }


	//fmt.Println(mybook.Current().Display())

	//   for {
	//       sen := mybook.Current()
	//
	//       b := make([]byte, 1)
	//       _, err = os.Stdin.Read(b)
	//       if err != nil {
	//           panic(err)
	//       }
	// //      display.Clear()
	//
	//       if strings.Compare(string(b), "\x03") == 0 {
	//           break
	//       } else if strings.Compare(string(b), "\x7f") == 0 {
	//           sen.Current.RemoveLast()
	//       } else if strings.Compare(string(b), "\r") == 0 {
	//           sen = mybook.Next();
	//       } else {
	//               sen.Current.Add(book.Letter{Char: b[0], Ignore: false})
	//       }
	//       if sen.Complete()  {
	//           sen = mybook.Next();
	//       }
	//
	//       //fmt.Printf("%q", b)
	//
	//
	//  //     fmt.Println(sen.Display())
	//   }
}
