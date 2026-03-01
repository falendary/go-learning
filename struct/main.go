package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//Capital letter = exported
//Lowercase letter = private

// Name → exported (public outside the package)
// name → unexported (private to the package)

// Declaration of Struct
type User struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Email     string   `json:"email"`
	Interests []string `json:"interests"`
}

// Method for User struct

func (u User) isAdult() bool {
	return u.Age >= 18
}

func (u User) Greet() string {
	return "Hello, " + u.Name
}

// Example with INT Concat
// Itoa = Integer to ASCII
func (u User) GreetWithAge() string {
	return "Hello, " + u.Name + "Age, " + strconv.Itoa(u.Age)
}

func (u User) GreetWithAgeSimple() string {
	return fmt.Sprintf("Hello fmt.Sprintf, %v Age, %v", u.Name, u.Age)
}

func (u User) listInterests() string {
	return strings.Join(u.Interests, ", ")
}

func main() {

	fmt.Println("Hello world")

	u := User{Name: "Serge", Age: 29, Email: "sz@example.com", Interests: []string{"Art", "Cinema", "Books"}}

	jsonData, _ := json.Marshal(u)
	fmt.Println("\nJSON:", string(jsonData))

	isAdult := u.isAdult()
	fmt.Println("Is Adult?", isAdult)

	fmt.Println(u.Greet())
	fmt.Println(u.GreetWithAge())
	fmt.Println(u.GreetWithAgeSimple())
	fmt.Println(u.listInterests())
}
