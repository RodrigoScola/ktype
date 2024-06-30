package book

import "github.com/RodrigoScola/ktype/pkg/sessions"


type Book struct {
    sentenceIndex int
    Sentences []Sentence `json:"sentences"`
    Session sessions.Session
}

func (b * Book) Add(sentence Sentence) {

    b.Sentences = append(b.Sentences, sentence)
}

func (b *Book) Current() *Sentence{
    return &b.Sentences[b.sentenceIndex]
}
func (b *Book) Next() *Sentence{
    b.sentenceIndex += 1
    return b.Current()
}



func NewBook(sentences []Sentence) *Book {
    return &Book{
        Sentences: sentences,
        sentenceIndex: 0,
    }
}
