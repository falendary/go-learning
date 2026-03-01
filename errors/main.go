package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {

	if b == 0 {
		return 0, fmt.Errorf("divide %w", errors.New("division by zero"))
	}

	return a / b, nil

}

func main() {
	// In Go: no try/catch. You check `err` after each call.
	// Think "handle or return".

	if result, err := divide(10, 2); err != nil {
		fmt.Println("Divide failed: ", err)
	} else {
		fmt.Println("Divide ok:", result)
	}

	if result, err := divide(10, 0); err != nil {
		fmt.Println("Divide failed: ", err)
	} else {
		fmt.Println("Divide ok:", result)
	}

}
