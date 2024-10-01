package main

import (
	"fmt"
	"slices"
)

func main() {
	letters := []string{"a", "b", "c", "d", "e", "f"}

	// chunk the slice into groups of 2
	for s := range slices.Chunk(letters, 2) {
		fmt.Println(s)
	}

	// collect the chunks into a slice
	s2 := slices.Collect(slices.Chunk(letters, 2))
	fmt.Println(s2)

	// collect the values back into a slice
	s3 := slices.Collect(slices.Values(letters))
	fmt.Println(s3)
}
