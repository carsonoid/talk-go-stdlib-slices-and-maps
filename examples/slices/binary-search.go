package main

import (
	"cmp"
	"fmt"
	"slices"
)

type user struct {
	name string
	age  int
}

func main() {
	letters := []string{"a", "b", "c", "d", "e"}
	i, ok := slices.BinarySearch(letters, "b")
	fmt.Println(i, ok) // 1 true

	users := []user{{"Tami", 25}, {"Raul", 30}, {"Tony", 35}}
	slices.SortFunc(users, cmpUser)
	slices.BinarySearchFunc(users, user{name: "Raul"}, cmpUser)
}

func cmpUser(a, b user) int {
	return cmp.Compare(a.name, b.name)
}
