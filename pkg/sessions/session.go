package sessions

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/RodrigoScola/ktype/pkg/book"
)

type SessionEntry struct {
	startAt  string
	endedAt  string
	Sentence []book.Sentence
}


type Session struct {
	Name    string `json:"name"`
	File    []string
    Entries []SessionEntry `json:"entries"`
}

func Exists(filename string) (*bool, error) {
	items, err := os.ReadDir(filepath.Join(".", "data", "sessions"))
	if err != nil {
		return nil, err
	}

	found := false
	for i := range items {
		if strings.Compare(items[i].Name(), filename) == 0 {
			found = true
			break
		}
	}

	return &found, nil
}

func GetSession(name string) (*Session, error) {
	var session Session

	items, err := os.ReadDir(filepath.Join(".", "data", "sessions"))
	if err != nil {
		return nil, err
	}

	for i := range items {
		if strings.Compare(items[i].Name(), name) == 0 {
			file, err := os.ReadFile(filepath.Join(".", "data", "sessions", name))

			if err != nil {
				return nil, err
			}
			json.Unmarshal(file, &session)

			break
		}
	}

	return &session, nil
}

func NewSession(name string, file []string) *Session {
	return &Session{Name: name, Entries: []SessionEntry{}, File: file}
}

func Save(sess *Session) (*Session, error) {
	bts, err := json.Marshal(sess)
	if err != nil {
		return nil, err
	}

	new_file := filepath.Join(".", "data", "sessions", sess.Name)
	err = os.WriteFile(new_file+".json", bts, 0644)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func Setup() (*Session, error) {
	new_file := filepath.Join(".", "data", "sessions")
	fmt.Println(new_file)

	err := os.MkdirAll(new_file, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return nil, nil
}



