package readers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func ReadByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)
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

func ReadByWord(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)
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

func ReadByChar(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)
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
func ReadSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)
	n, err := f.Read(buffer)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buffer[0:n]
}
