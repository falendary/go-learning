package main

import (
	"fmt"
	"slices"
	"sort"
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
