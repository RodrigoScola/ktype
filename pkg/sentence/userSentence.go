package sentence



type UserSentence struct {
    Correct Sentence
    Current Sentence
}


func ToLetters(str string) []Letter {
    letters := make([]Letter,0)

    for _, i := range str {
        letters = append(letters, Letter{
            char: i,
        })
    }
    return letters
}

func Display() {
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
