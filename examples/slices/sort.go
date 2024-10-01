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

	fmt.Println(slices.IsSorted(s))

	slices.Sort(s)
	fmt.Println(s)

	fmt.Println(slices.IsSorted(s))
}
