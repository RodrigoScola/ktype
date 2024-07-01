package book


type Book struct {
	sentenceIndex int
	Sentences     []Sentence `json:"sentences"`
Name string `json:"name"`
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
    filled := []*Sentence{}
    for i := range b.Sentences {
        if len (b.Sentences[i].Current.Letters) == 0 {
            break
        }
        filled = append(filled, &b.Sentences[i])
    }
    return filled
}
func (b *Book) Unfilled() []*Sentence {
    unfilled := []*Sentence{}
    for i := range b.Sentences {
        if len (b.Sentences[i].Current.Letters) == 0 {
            unfilled = append(unfilled, &b.Sentences[i])
        }
    }
    return unfilled
}

func NewBook(name string ,sentences []Sentence) *Book {
	return &Book{
        Name: name,
		Sentences:     sentences,
		sentenceIndex: 0,
	}
}
