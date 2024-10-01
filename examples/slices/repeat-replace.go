package main

import (
	"fmt"
	"slices"
)

type user struct {
	name string
	age  int
}

func main() {
	s := []string{}

	s = slices.Repeat([]string{"a", "b", "c"}, 3)
	fmt.Println(s)

	s = slices.Replace(s, 0, 3, "x", "y", "z")
	fmt.Println(s)
}
