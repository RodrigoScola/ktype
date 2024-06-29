package book

import (
	"strings"
)
type Letter struct {
    Char byte;
    Ignore bool
}

type Sentence struct {
    Letters []Letter
}
func (s *Sentence) Length() int {
    curr := 0;
    for i := 0; i < len(s.Letters); i++ {
        if  s.Letters[i].Ignore == true {
            continue
        }
        curr++
    }
    return curr
}

func (s* Sentence) GetLast() *Letter {
    for i := len(s.Letters) -1; i >= 0; i-- {
        if s.Letters[i].Ignore {
            continue
        }
        return &s.Letters[i]
    }
    return nil
}

func (s *Sentence) RemoveLast() {
    last := s.GetLast()
    if last != nil {
        last.Ignore = true
    }
}
func (s* Sentence) Remove(ind int) {
    curr :=0;
    for i := 0; i < len(s.Letters); i++ {
        if  s.Letters[i].Ignore {
            continue
        } else if curr == ind  {
            s.Letters[curr].Ignore = true;

            continue
        }
        curr++
    }

}


func (s *Sentence) Add(letter Letter) *Sentence {
    s.Letters = append(s.Letters, letter)
    return s
}

func (s *Sentence) String() string {
    var sb strings.Builder

    for _, c := range s.Letters {
        if c.Ignore == false {
            sb.WriteByte(c.Char)
        }
    }
    return sb.String()
}
