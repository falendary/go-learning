package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Create("multi.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// MultiWriter writes to stdout and the file at the same time
	w := io.MultiWriter(os.Stdout, f)

	fmt.Fprintln(w, "This line goes to terminal and file")
}
