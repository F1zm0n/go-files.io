package files

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

func (f *FileRead) ReadByLine() error {
	r := bufio.NewReader(f.File)
	var counter = 0
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			if counter == 1 {
				break
			}
			counter++
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		fmt.Print(line)
	}
	return nil
}

func (f *FileRead) ReadByWord() error {
	r := bufio.NewReader(f.File)
	var counter = 0
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			if counter == 1 {
				break
			}
			counter++
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}
		r := regexp.MustCompile("[^\\s]+")
		words := r.FindAllString(line, -1)

		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
		}
	}
	return nil
}

func (f *FileRead) ReadByChar() error {
	r := bufio.NewReader(f.File)
	var counter = 0
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			if counter == 1 {
				break
			}
			counter++
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		for _, x := range line {
			fmt.Print(string(x))
		}

	}
	return nil
}
func (f *FileRead) ReadSize(size int) []byte {
	buffer := make([]byte, size)
	n, err := f.File.Read(buffer)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return nil
	}
	return buffer[0:n]
}
