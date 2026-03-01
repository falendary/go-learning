package tests

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type ValidationError struct {
	Field string
	Msg   string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s: %s", e.Field, e.Msg)
}

func Sum(nums []int) int {

	total := 0

	for _, val := range nums {
		total = total + val
	}

	return total
}

func Divide(a, b int) (int, error) {

	if b == 0 {
		return 0, fmt.Errorf("divide %w", ValidationError{Field: "b", Msg: "must not be zero"})
	}

	return a / b, nil
}

// Capital Case -> being "public" and exported outside file
func FindUserEmail(id string) (string, error) {

	if id == "" {
		return "", ValidationError{Field: "id", Msg: "must not be empty"}
	}

	if id != "42" {
		return "", fmt.Errorf("findUserEmail(%s): %w", id, ErrNotFound)
	}

	return "sz@example.com", nil
}
