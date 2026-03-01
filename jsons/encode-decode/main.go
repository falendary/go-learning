package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
User demonstrates:

- JSON tags
- omitempty
- exported fields only
*/

type User struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Email     string   `json:"email,omitempty"`
	Interests []string `json:"interests"`
}

func main() {

	// =========================
	// Struct → JSON (Encode)
	// =========================

	user := User{
		Name:      "Serge",
		Age:       29,
		Email:     "",
		Interests: []string{"Art", "Cinema"},
	}

	fmt.Println("Struct → JSON (Encode):")

	/*
		NewEncoder takes an io.Writer.
		os.Stdout is the terminal output.
	*/
	encoder := json.NewEncoder(os.Stdout)

	// Pretty formatting
	encoder.SetIndent("", "  ")

	err := encoder.Encode(user)
	if err != nil {
		panic(err)
	}

	// =========================
	// JSON → Struct (Decode)
	// =========================

	raw := `{
		"name": "Anna",
		"age": 34,
		"email": "anna@example.com",
		"interests": ["Books", "Travel"]
	}`

	var u2 User

	/*
		NewDecoder takes an io.Reader.
		strings.NewReader turns string into Reader.
	*/
	decoder := json.NewDecoder(strings.NewReader(raw))

	err = decoder.Decode(&u2)
	if err != nil {
		fmt.Println("Decode error:", err)
		return
	}

	fmt.Println("\nJSON → Struct (Decode):")
	fmt.Printf("%+v\n", u2)

	// =========================
	// JSON array → slice
	// =========================

	rawArray := `[
		{"name":"A","age":20,"interests":[]},
		{"name":"B","age":25,"interests":["Music"]}
	]`

	var users []User

	arrayDecoder := json.NewDecoder(strings.NewReader(rawArray))

	err = arrayDecoder.Decode(&users)
	if err != nil {
		fmt.Println("Decode array error:", err)
		return
	}

	fmt.Println("\nJSON array → slice:")
	fmt.Println(users)
}
