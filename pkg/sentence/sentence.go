package sentence

import "strings"
type Letter struct {
    char rune;
}

type Sentence struct {
    Letters []Letter
}



func (s *Sentence) String() string {
    var sb strings.Builder

    for _, c := range s.Letters {
        sb.WriteRune(c.char)
    }
    return sb.String()
}
