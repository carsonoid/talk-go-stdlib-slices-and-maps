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
	s := []string{"a", "b", "c", "c", "a", "a"}
	s = slices.Compact(s)
	fmt.Println(s)

	users := []user{
		{"Tami", 25},
		{"Tami", 25},
		{"Raul", 30},
		{"Tony", 35},
		{"Tami", 25},
	}

	fmt.Println(slices.CompactFunc(users, userByName))
}

func userByName(a, b user) bool {
	return a.name == b.name
}
