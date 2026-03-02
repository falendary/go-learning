package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	// =========================
	//  Write file
	// =========================

	content := []byte("Hello file\nSecond line\n")

	err := os.WriteFile("example.txt", content, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("File written: example.txt")

	// =========================
	// Read entire file
	// =========================

	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("\nReadFile output:")
	fmt.Println(string(data))

	// =========================
	// Append to file
	// =========================

	f, err := os.OpenFile("example.txt",
		os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString("Appended line\n")
	if err != nil {
		panic(err)
	}
	f.Close()

	fmt.Println("\nAppended one line")

	// =========================
	// Read line by line
	// =========================

	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("\nReading line by line:")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("Line:", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// =========================
	// In-memory file (bytes.Buffer)
	// =========================

	fmt.Println("\nIn-memory file example:")

	var buf bytes.Buffer

	// bytes.Buffer implements io.Writer
	buf.WriteString("Hello in memory\n")
	buf.WriteString("Another line\n")

	fmt.Println("Buffer content:")
	fmt.Println(buf.String())

	// You can also read from it
	reader := bytes.NewReader(buf.Bytes())

	fmt.Println("Reading via io.Reader:")
	io.Copy(os.Stdout, reader)
}
