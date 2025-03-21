package file

import (
	"fmt"
	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, content []byte) (*File, error) {
	fileID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &File{
		ID:   fileID,
		Name: filename,
		Data: content,
	}, nil
}

func (f *File) String() string {
	return fmt.Sprintf("ID: %s, Name: %v", f.ID, f.Name)
}
