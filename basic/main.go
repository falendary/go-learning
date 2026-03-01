package main

import "fmt"

//import "encoding/json"

func main() {

	var n int
	var text string
	var is_active bool

	// Zero Values
	fmt.Println("Zero Values:")
	fmt.Println("int:", n)
	fmt.Println("string:", text)
	fmt.Println("bool:", is_active)

	// Val vs Pointer
	x := 10
	fmt.Println("x initial:", x)
	x++
	fmt.Println("x increment:", x)

}
