package main

import (
	"bytes"
	"fmt"
)

func main() {
	// bytes.Buffer is an in-memory writer
	var buf bytes.Buffer

	// WriteString writes into the buffer
	buf.WriteString("Line one\n")
	buf.WriteString("Line two\n")

	fmt.Println(buf.String())
}
