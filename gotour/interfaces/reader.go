package main

import (
	"fmt"
	"io"
	"golang.org/x/tour/reader"
)

type Reader interface {
	Read(b []byte)
}

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(b []byte) (n int, err error) {
	n = 0
	for i := range b {
		b[i] = 65
		n++
		fmt.Println("A, ")
	} /*
		for {
			fmt.Println("A, ")
			n++
		}*/
	//	fmt.Println("hey!", b[0])
	fmt.Println(n)
	err = io.EOF
	return
}

func main() {
	reader.Validate(MyReader{})
}