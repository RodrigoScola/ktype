package filesessions

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/RodrigoScola/ktype/pkg/book"
)

func getSessionsPath() string {
	return path.Join(".", "data", "sessions")
}

func Save(book *book.Book) (*book.Book, error) {
	mark, err := json.Marshal(book)
	if err != nil {
		return nil, err
	}
	pa := getSessionsPath() + "/" + book.Name + ".json"

	err = os.WriteFile(pa, mark, 0644)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func GetSessionNames() ([]string, error) {
	files, err := os.ReadDir(getSessionsPath())

	names := make([]string, len(files))
	if err != nil {
		return []string{}, err
	}

	for i, file := range files {
		names[i] = file.Name()
	}
	return names, nil

}
func GetSession(name string) (*book.Book, error) {
	var filesession book.Book
	files, err := os.ReadDir(getSessionsPath())
	if err != nil {
		return nil, err
	}
	var sess_file fs.DirEntry
	for _, file := range files {
		if strings.Compare(name, file.Name()) == 0 {
			sess_file = file
			break
		}
	}
	if sess_file == nil {
		return nil, errors.New("file not found")
	}
	pa := getSessionsPath() + "/" + name
	f, err := os.ReadFile(pa)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(f, &filesession)
	if err != nil {
		return nil, err
	}

	return &filesession, nil

}
