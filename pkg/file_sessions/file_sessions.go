package filesessions

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/RodrigoScola/ktype/pkg/book"
)

type FileSession struct {
	Entries *[]book.Sentence `json:"book"`
	File    *[]string
	Name    *string
    Book *book.Book
}

func New(
	name *string, file *[]string, entries *[]book.Sentence,
) *FileSession {
	return &FileSession{
		Name:    name,
		File:    file,
		Entries: entries,
	}
}

func (f *FileSession) Save() error {
    f.Entries = &f.Book.Sentences
	// Modify the Save function to be a method of FileSession
	mark, err := json.Marshal(f)

    fmt.Println(f.Entries)
	if err != nil {
		return err
	}
    fmt.Println(f.Name, *f.Name)

	pa := getSessionsPath() + "/" + *f.Name + ".json"

	err = os.WriteFile(pa, mark, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getSessionsPath() string {
	return path.Join(".", "data", "sessions")
}
func Save(file *FileSession) (*FileSession, error) {
	mark, err := json.Marshal(file)
	if err != nil {
		return nil, err
	}
   fmt.Println(file, "this is the file")
	 pa := getSessionsPath() + "/" + *file.Name + ".json"
	 fmt.Println(pa)

	err = os.WriteFile(pa, mark, 0644)
	 if err != nil {
	 	return nil, err
	 }

	return file, nil
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
func GetSession(name string) (*FileSession, error) {

    var filesession FileSession
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
    if sess_file == nil  {
        return nil, errors.New("file not found")
    }
    f , err := os.ReadFile(getSessionsPath() + "/"+ name)
    if err != nil {
        return nil, err
    }


    err = json.Unmarshal(f, &filesession)
    if err != nil {
        return nil, err
    }

    return &filesession, nil

}
