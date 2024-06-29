package book


type Book struct {
    sentenceIndex int
    Sentences []UserSentence
}

func (b * Book) Add(sentence UserSentence) {

    b.Sentences = append(b.Sentences, sentence)
}

func (b *Book) Current() *UserSentence{
    return &b.Sentences[b.sentenceIndex]
}
func (b *Book) Next() *UserSentence{
    b.sentenceIndex += 1
    return b.Current()
}



func NewBook() Book {
    return Book{
        Sentences: make([]UserSentence, 0),
        sentenceIndex: 0,
    }
}
