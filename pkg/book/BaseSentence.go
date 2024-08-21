package book

import (
	"strings"
	"time"
)
type Letter struct {
    Char byte `json:"char"`
    CreatedAt time.Time `json:"createdAt"`
    Ignore bool `json:"ignore"`
}

type BaseSentence struct {
Letters []Letter `json:"letters"`
}



func (s *BaseSentence) Length() int {
    curr := 0;
    for i := 0; i < len(s.Letters); i++ {
        if  s.Letters[i].Ignore == true {
            continue
        }
        curr++
    }
    return curr
}
func (s* BaseSentence) GetFirst() *Letter {
    for i :=0; i < len(s.Letters);i++ {
        if s.Letters[i].Ignore {
            continue
        }
        return &s.Letters[i]
    }
    return nil
}

func (s* BaseSentence) GetLast() *Letter {
    for i := len(s.Letters) -1; i >= 0; i-- {
        if s.Letters[i].Ignore {
            continue
        }
        return &s.Letters[i]
    }
    return nil
}

func (s *BaseSentence) RemoveLast() {
    last := s.GetLast()
    if last != nil {
        last.Ignore = true
    }
}
func (s* BaseSentence) Remove(ind int) {
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


func (s *BaseSentence) Add(letter Letter) *BaseSentence {
    s.Letters = append(s.Letters, letter)
    return s
}

func (s *BaseSentence) String() string {
    var sb strings.Builder

    for _, c := range s.Letters {
        if c.Ignore == false {
            sb.WriteByte(c.Char)
        }
    }
    return sb.String()
}
