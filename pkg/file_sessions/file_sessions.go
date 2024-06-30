package filesessions

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/RodrigoScola/ktype/pkg/book"
)

type FileSession struct {
    Entries *[]book.Sentence `json:"book"`
    File *[]string
    Name *string
}

func New(
    name *string, file *[]string, entries *[]book.Sentence,
) *FileSession {
    return &FileSession{
        Name: name,
        File: file,
        Entries: entries,
    }
}

func getSessionsPath() string {
    return path.Join(".", "data", "sessions")
}
func Save(file *FileSession) (*FileSession, error) {
    mark, err := json.Marshal(file)
    if err != nil {
        return nil, err
    }
    pa:= getSessionsPath() + "/" + *file.Name + ".json"
    fmt.Println(pa)

    err = os.WriteFile( pa,mark,0644)
    if err != nil {
        return nil, err
    }

    return file, nil
}
