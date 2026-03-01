package main

import "fmt"

/*
==============================
WITHOUT GENERICS
==============================

You must write a separate function
for each type.
*/

// FilterInts keeps ints that satisfy keep()
func FilterInts(nums []int, keep func(int) bool) []int {
	out := []int{}
	for _, n := range nums {
		if keep(n) {
			out = append(out, n)
		}
	}
	return out
}

// FilterStrings keeps strings that satisfy keep()
func FilterStrings(items []string, keep func(string) bool) []string {
	out := []string{}
	for _, s := range items {
		if keep(s) {
			out = append(out, s)
		}
	}
	return out
}

/*
Problem:

- Logic is identical
- Code is duplicated
- If you change implementation,
  you must change it in multiple places
*/

/*
==============================
WITH GENERICS
==============================

One function.
Works for any type.
*/

func Filter[T any](items []T, keep func(T) bool) []T {
	out := []T{}
	for _, item := range items {
		if keep(item) {
			out = append(out, item)
		}
	}
	return out
}

func main() {

	// ---- Without generics ----
	nums := []int{1, 2, 3, 4, 5}
	evensOld := FilterInts(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Without generics (ints):", evensOld)

	names := []string{"Serge", "Anna", "Max"}
	longNamesOld := FilterStrings(names, func(s string) bool {
		return len(s) >= 4
	})
	fmt.Println("Without generics (strings):", longNamesOld)

	// ---- With generics ----
	evens := Filter(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("With generics (ints):", evens)

	longNames := Filter(names, func(s string) bool {
		return len(s) >= 4
	})
	fmt.Println("With generics (strings):", longNames)
}
