package main

import (
	"fmt"
	"io"
	"strings"
)

/*
IO in Go is built around two core interfaces:

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

Everything else builds on top of these.
*/

func main() {
	// strings.NewReader creates an io.Reader from a string
	r := strings.NewReader("Hello IO\nSecond line\n")

	// io.ReadAll reads everything from a Reader into memory
	data, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
