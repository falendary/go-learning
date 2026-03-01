package main

import (
	"errors"
	"fmt"
	"github.com/falendary/go-learning/errors/shared"
)

func main() {

	err := loadConfig("/etc/app/config.json")
	if err != nil {
		fmt.Println("\nloadConfig error:", err)
		fmt.Println("-> you can still decide what to do at the top level")
	}

}

func loadConfig(path string) error {

	if path == "" {
		return fmt.Errorf("loadConfig: %w", shared.ValidationError{Field: "path", Msg: "empty"})
	}

	// Imagine an underlying error from IO, parsing, etc
	underlying := errors.New("permission denied")

	return fmt.Errorf("loadConfig(%s): %w", path, underlying)

}
