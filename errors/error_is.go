package main

import (
	"errors"
	"fmt"
)

func main() {

	email, err := getUserEmail("100")

	if err != nil {

		fmt.Println("\ngetUserEmail error", err)

		if errors.Is(err, ErrNotFound) {
			fmt.Println("-> detected ErrNotFound via errors.Is")
		}

	} else {
		fmt.Println("getUserEmail() -> email", email)
	}

	email, err = getUserEmail("")

	if err != nil {

		fmt.Println("\ngetUserEmail error", err)

		var ve ValidationError

		if errors.As(err, &ve) {
			fmt.Println("-> validation details, field=%s, message=%s", ve.Field, ve.Msg)
		}

	} else {
		fmt.Println("getUserEmail() -> email", email)
	}

}

var ErrNotFound = errors.New("not found")

type ValidationError struct {
	Field string
	Msg   string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s: %s", e.Field, e.Msg)
}

func getUserEmail(id string) (string, error) {

	if id == "" {
		return "", ValidationError{Field: "id", Msg: "must not be empty"}
	}

	if id != "42" {
		//%v → formats as text only
		//%w → wraps the error structurally
		return "", fmt.Errorf("getUserEmail(%s): %w", id, ErrNotFound)
	}

	return "sz@example.com", nil

}
