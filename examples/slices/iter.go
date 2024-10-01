package main

import (
	"fmt"
	"slices"
)

func main() {
	letters := []string{"a", "b", "c"}

	// iterate over the slice with the slices package
	for _, s := range letters {
		fmt.Println(s)
	}

	// iterate using `iter.Seq2` (value and index)
	for _, s := range slices.All(letters) {
		fmt.Println(s)
	}

	// iterate using `iter.Seq` (value only)
	for s := range slices.Values(letters) {
		fmt.Println(s)
	}
}
