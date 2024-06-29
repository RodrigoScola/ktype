package book

import (
	"strings"

	"github.com/RodrigoScola/ktype/pkg/display"
)

type UserSentence struct {
	Correct Sentence
	Current Sentence
}

func ToLetters(str string) []Letter {
	letters := make([]Letter, 0)

	for _, i := range []byte(str) {
		letters = append(letters, Letter{
			Char:   i,
			Ignore: false,
		})
	}
	return letters
}

func (s *UserSentence) Complete() bool {
    if strings.Compare(s.Correct.String(), s.Current.String()) == 0 {
        return true
    }
	return false
}

func (s *UserSentence) Display() string {
	ind := 0
    isWrong := false
	var sb strings.Builder
	for _, char := range s.Current.Letters {
		if char.Ignore == true {
			continue
		}
		if isWrong == false && char.Char == s.Correct.Letters[ind].Char {
            sb.WriteString(display.ColorsVar.FgYellow)
            sb.WriteByte(char.Char)
            sb.WriteString(display.ColorsVar.Reset)
		ind++
		} else {
            isWrong = true
            sb.WriteString(display.ColorsVar.FgRed)
            sb.WriteByte(char.Char)
            sb.WriteString(display.ColorsVar.Reset)
        }

	}
    for i := ind  ; i < s.Correct.Length() ; i++ {
            sb.WriteString(display.ColorsVar.FgWhite)
            sb.WriteByte(s.Correct.Letters[i].Char)
            sb.WriteString(display.ColorsVar.Reset)
    }

	return sb.String()
}

func NewUserSentence(sen string) UserSentence {

	return UserSentence{
		Correct: Sentence{
			Letters: ToLetters(sen),
		},
		Current: Sentence{
			Letters: make([]Letter, 0),
		},
	}
}
