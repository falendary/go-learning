package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// io.Copy copies bytes from Reader to Writer
	src := strings.NewReader("Copy demo\n")

	fmt.Println("Copying to stdout")
	_, err := io.Copy(os.Stdout, src)
	if err != nil {
		panic(err)
	}
}
