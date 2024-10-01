package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c", "d", "e"}
	fmt.Println(s)

	// index
	fmt.Println(s[0])

	// slice a slice
	fmt.Println(s[1:3])

	// add an element
	s = append(s, "f")
	fmt.Println(s)
}
