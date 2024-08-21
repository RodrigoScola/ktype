package book

import (
	"fmt"

	"github.com/RodrigoScola/ktype/pkg/formatting"
)

type Book struct {
	sentenceIndex int
	Sentences     []Sentence `json:"sentences"`
	Name          string     `json:"name"`
}

func (b *Book) Add(sentence Sentence) {
	b.Sentences = append(b.Sentences, sentence)
}

func (b *Book) Current() *Sentence {
	return &b.Sentences[b.sentenceIndex]
}
func (b *Book) Next() *Sentence {
	b.sentenceIndex += 1
	return b.Current()
}

func (b *Book) FilledSentences() []*Sentence {
	filled := make([]*Sentence, 0)
        fmt.Println("he y")
	for _, v := range b.Sentences {
		if len(v.Current.Letters) == 0 {
            fmt.Println("v.Current.Letters", v.Current.Letters)
			continue
		}
		v.Correct = formatting.FormatAll(v.Correct)
		filled = append(filled, &v)
	}
	return filled
}
func (b *Book) Unfilled() []*Sentence {
	unfilled := []*Sentence{}
	for _, v := range b.Sentences {
		if len(v.Current.Letters) > 0 {
			continue
		}
		v.Correct = formatting.FormatAll(v.Correct)
		unfilled = append(unfilled, &v)
	}
	return unfilled
}

func NewBook(name string, sentences []Sentence) *Book {
	return &Book{
		Name:          name,
		Sentences:     sentences,
		sentenceIndex: 0,
	}
}
