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
	s := []string{"b", "c", "a", "e", "d"}
	fmt.Println(cap(s))

	s = slices.Grow(s, 10)
	fmt.Println(cap(s))

	s = slices.Clip(s)
	fmt.Println(cap(s))
}
