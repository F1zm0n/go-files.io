package readers

import "os"

type FileRead struct {
	File *os.File
}

func NewFileReader(file *os.File) *FileRead {
	return &FileRead{file}
}

func OpenFileReader(filename string) (*FileRead, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return &FileRead{f}, nil
}
