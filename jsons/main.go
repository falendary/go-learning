package main

import (
	"encoding/json"
	"fmt"
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
	// Struct → JSON
	// =========================

	user := User{
		Name:      "Serge",
		Age:       29,
		Email:     "",
		Interests: []string{"Art", "Cinema"},
	}

	// Marshal converts Go struct to JSON bytes
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Struct → JSON:")
	fmt.Println(string(jsonBytes))
	// Notice: email is omitted because of `omitempty`

	// Pretty print
	pretty, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("\nPretty JSON:")
	fmt.Println(string(pretty))

	// =========================
	// JSON → Struct
	// =========================

	raw := `{
		"name": "Anna",
		"age": 34,
		"email": "anna@example.com",
		"interests": ["Books", "Travel"]
	}`

	var u2 User

	err = json.Unmarshal([]byte(raw), &u2)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}

	fmt.Println("\nJSON → Struct:")
	fmt.Printf("%+v\n", u2)

	// =========================
	// JSON array → slice
	// =========================

	rawArray := `[
		{"name":"A","age":20,"interests":[]},
		{"name":"B","age":25,"interests":["Music"]}
	]`

	var users []User

	err = json.Unmarshal([]byte(rawArray), &users)
	if err != nil {
		fmt.Println("Unmarshal array error:", err)
		return
	}

	fmt.Println("\nJSON array → slice:")
	fmt.Println(users)

}
