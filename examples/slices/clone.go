package main

import (
	"fmt"
	"slices"
)

func before() {
	s1 := []string{"a", "b", "c"}
	s2 := append([]string(nil), s1...)
	s1[0] = "x"
	fmt.Println(s1)
	fmt.Println(s2)
}

func main() {
	s1 := []string{"a", "b", "c"}
	s2 := slices.Clone(s1)
	s1[0] = "x"
	fmt.Println(s1)
	fmt.Println(s2)
}
