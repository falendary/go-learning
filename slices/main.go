package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func main() {

	// create slice
	nums := []int{1, 10, 5, 15}
	fmt.Println("nums: ", nums)

	// Create with make
	names := make([]string, 0, 4)
	names = append(names, "Serge", "Sara", "Luna")
	fmt.Println("names: ", names, " len: ", len(names), " cap: ", cap(names))

	// Append item
	nums = append(nums, 20)
	fmt.Println("append 20: ", nums)

	// Remove item by index
	nums = slices.Delete(nums, 2, 2+1) // remove element at index 2
	fmt.Println("slices.Delete(nums, i, i+1): ", nums)
	nums = append(nums, 5)

	nums = removeAt(nums, 3)
	fmt.Println("custom implementation: removeAt(nums, 3)", nums)
	nums = append(nums, 20)

	// Remove item by value
	names = removeValue(names, "Sara")
	fmt.Println("custom implementation: removeValue(names, \"Sara\")", names)

	// Sort
	slices.Sort(nums)
	fmt.Println("sorted nums: ", nums)

	// Sort reverse
	slices.Reverse(nums)
	fmt.Println("sorted revers nums: ", nums)

	// Sort names

	sort.Strings(names)
	fmt.Println("sorted names", names)

	// Filters
	evens := filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("even nums:", evens)

	// Join slices

	first := []int{100, 200}
	second := []int{300, 400}

	//Why is it needed?
	//
	//The append function looks like this:
	//
	//append(slice []T, elems ...T) []T
	//
	//Notice ...T.
	//
	//	That means:
	//
	//The second parameter is variadic
	//
	//It expects individual values, not a slice

	first = append(first, second...) // The ... is called the variadic expansion operator.
	fmt.Println("Joined slices", first)

	//Using slices.Concat
	third := []int{500, 600}
	forth := []int{700, 800}

	joined := slices.Concat(third, forth)
	fmt.Println("Joined using slices.Concat", joined)

	// Find item
	// if initialization; condition {
	if idx, ok := findIndex(names, "Kate"); ok {
		fmt.Println("Name Kate is found", idx)
	} else {
		fmt.Println("Kate is not found")
	}

	// Replace item
	names = replaceFirst(names, "Serge", "Sergei")
	fmt.Println("Replace name", names)

	// Join Strings
	fmt.Println("name strings", strings.Join(names, ", "))
}

// removeAt removes element at index i, preserving order
func removeAt[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	// join to slices, excluding i
	// [1, 10, 5, 15, 20], exclude index 2
	// [1, 10] + [15, 20] => [1 ,10, 15, 20]

	return append(s[:i], s[i+1:]...)
}

func removeValue[T comparable](s []T, v T) []T {

	idx, ok := findIndex(s, v)

	if ok {
		return removeAt(s, idx)
	}

	return s
}

func findIndex[T comparable](s []T, target T) (int, bool) {

	for i, v := range s {
		if v == target {
			return i, true
		}
	}

	return -1, false

}

// filter returns a new slice containing items for which keep(item) == true
func filter[T any](s []T, keep func(T) bool) []T {

	out := make([]T, 0, len(s))

	for _, val := range s {
		if keep(val) {
			out = append(out, val)
		}
	}

	return out
}

func replaceFirst[T comparable](s []T, old T, new T) []T {

	if idx, ok := findIndex(s, old); ok {
		s[idx] = new
	}

	return s

}
