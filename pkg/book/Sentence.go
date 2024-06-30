package book

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Sentence struct {
	Correct BaseSentence
	Current BaseSentence
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

func (s *Sentence) Complete() bool {
    if strings.Compare(s.Correct.String(), s.Current.String()) == 0 { return true
    }
	return false
}

func (s *Sentence) Display(primary lipgloss.Style, secondary lipgloss.Style) string {
	ind := 0
    isWrong := false
	var sb strings.Builder
	for _, char := range s.Current.Letters {
		if char.Ignore == true {
			continue
		}
		if isWrong == false && char.Char == s.Correct.Letters[ind].Char {
            sb.WriteString(primary.Render(string(char.Char)))
            ind++
		} else {
            isWrong = true

            sb.WriteString(secondary.Render(string(char.Char)))
        }

	}
    for i := ind  ; i < s.Correct.Length() ; i++ {
            sb.WriteByte(s.Correct.Letters[i].Char)
    }

	return sb.String()
}

func NewUserSentence(sen string) Sentence {

	return Sentence{
		Correct: BaseSentence{
			Letters: ToLetters(sen),
		},
		Current: BaseSentence{
			Letters: make([]Letter, 0),
		},
	}
}
