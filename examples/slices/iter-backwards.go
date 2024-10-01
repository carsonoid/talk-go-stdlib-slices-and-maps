package main

import (
	"fmt"
	"slices"
)

func main() {
	letters := []string{"a", "b", "c"}

	// iterate in reverse
	for _, s := range slices.Backward(letters) {
		fmt.Println(s)
	}

	// use an iterator to sort the slice
	unsorted := []string{"c", "a", "b"}
	for _, s := range slices.Sorted(slices.Values(unsorted)) {
		fmt.Println(s)
	}
}
