package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	source := strings.NewReader("Tee demo\n")

	// mirror will receive a copy of everything we read from source
	var mirror bytes.Buffer

	// TeeReader reads from source and writes the same bytes into mirror
	tee := io.TeeReader(source, &mirror)

	// Copy tee output to stdout
	_, err := io.Copy(os.Stdout, tee)
	if err != nil {
		panic(err)
	}

	fmt.Println("Mirror buffer content")
	fmt.Print(mirror.String())
}
