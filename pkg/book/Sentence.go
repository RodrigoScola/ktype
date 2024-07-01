package book

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Sentence struct {
Correct string `json:"correct"`
Current BaseSentence `json:"base"`
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

func ToSentences(strs []string) []Sentence {
	sentences := make([]Sentence, len(strs))

	for i := range strs {
		sentences = append(sentences, Sentence{
			Correct: strs[i],
			Current: BaseSentence{Letters: []Letter{}},
		})
	}
	return sentences
}

func (s *Sentence) Complete() bool {
	if strings.Compare(s.Correct, s.Current.String()) == 0 {
		return true
	}
	return false
}

func (s *Sentence) Display(primary lipgloss.Style, secondary lipgloss.Style, text lipgloss.Style) string {
	 ind := 0
	 isWrong := false
	 var sb strings.Builder
	 for _, char := range s.Current.Letters {
	 	if char.Ignore == true {
	 		continue
	 	}
	 	if isWrong == false && char.Char == s.Correct[ind] {
	 		sb.WriteString(primary.Render(string(char.Char)))
	 		ind++
	 	} else {
	 		isWrong = true
	 		sb.WriteString(secondary.Render(string(char.Char)))
	 	}
	 }
	 for i := ind; i < len(s.Correct); i++ {
	 	sb.WriteString(text.Render( string(s.Correct[i])))
	 }

	return sb.String()
}

func NewUserSentence(sen string) Sentence {

	return Sentence{
		Correct: sen,
		Current: BaseSentence{
			Letters: make([]Letter, 0),
		},
	}
}
