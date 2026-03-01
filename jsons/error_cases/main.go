package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*
User demonstrates:

- JSON tags
- omitempty
- exported fields only (capital letter)
*/
type User struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Email     string   `json:"email,omitempty"`
	Interests []string `json:"interests"`
}

func main() {

	// =========================
	// Parse error example
	// =========================

	invalidJSON := `{"name": "Bad JSON", "age": }`

	var bad User
	err := json.Unmarshal([]byte(invalidJSON), &bad)

	fmt.Println("\nInvalid JSON example:")
	fmt.Println("Error:", err)

	// =========================
	// Unknown fields handling
	// =========================

	unknownFieldJSON := `{
		"name": "Extra",
		"age": 30,
		"unknown_field": "something",
		"interests": []
	}`

	// Default behavior: unknown fields are ignored
	var u3 User
	err = json.Unmarshal([]byte(unknownFieldJSON), &u3)
	fmt.Println("\nUnknown fields (default):")
	fmt.Println("Error:", err)
	fmt.Printf("%+v\n", u3)

	// Strict mode: disallow unknown fields
	fmt.Println("\nUnknown fields (strict mode):")

	decoder := json.NewDecoder(strings.NewReader(unknownFieldJSON))
	decoder.DisallowUnknownFields()

	var strictUser User
	err = decoder.Decode(&strictUser)
	fmt.Println("Error:", err)
}
