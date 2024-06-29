package main

import (
	"fmt"

	"github.com/RodrigoScola/ktype/pkg/sentence"
)


func main() {
    sen := sentence.NewUserSentence("this should be the correct thing")
    fmt.Println("hel")


    fmt.Println(sen.Correct.String())
}
