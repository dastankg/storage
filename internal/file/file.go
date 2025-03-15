package file

type File struct {
	Name string
	Data []byte
}

func NewFile() *File {
	return &File{}
}
