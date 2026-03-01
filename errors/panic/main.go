package main

import "fmt"

func main() {

	safeRun(func() {
		mightPanic(true)
	})

}

func mightPanic(boom bool) {

	if boom {
		// panic is a built-in function that stops normal execution immediately.
		// Panic is reserved for serious programmer mistakes
		// is similar to raise Exception
		// in go Errors are returned explicitly, panic is NOT normal control flow
		panic("something very very wrong")
	}
}

func safeRun(fn func()) {

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("recovered from panic", r)
		}

	}()
	fn()
}
