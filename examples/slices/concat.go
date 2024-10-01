package main

import (
	"fmt"
	"slices"
)

type user struct {
	name string
	age  int
}

func before() {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"d", "e", "f"}
	s3 := []string{"g", "h", "i"}
	all := append(s1, s2...)
	all = append(all, s3...)
	fmt.Println(all)
}

func main() {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"d", "e", "f"}
	s3 := []string{"g", "h", "i"}
	all := slices.Concat(s1, s2, s3)
	fmt.Println(all)
}
