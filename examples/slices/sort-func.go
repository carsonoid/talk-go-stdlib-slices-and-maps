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
	users := []user{{"Raul", 30}, {"Tami", 25}, {"Tony", 35}}

	fmt.Println(slices.IsSortedFunc(users, sortByAge))

	slices.SortFunc(users, sortByAge)
	fmt.Println(users)

	fmt.Println(slices.IsSortedFunc(users, sortByAge))
}

func sortByAge(i, j user) int {
	return cmp.Compare(i.age, j.age)
}
