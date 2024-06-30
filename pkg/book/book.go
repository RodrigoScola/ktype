package book


type Book struct {
    sentenceIndex int
    Sentences []Sentence
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



func NewBook() *Book {
    return &Book{
        Sentences: make([]Sentence, 0),
        sentenceIndex: 0,
    }
}
