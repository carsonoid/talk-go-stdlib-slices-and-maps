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
	s := []string{"a", "b", "f", "g"}
	s = slices.Insert(s, 2, "c", "d", "e")
	fmt.Println(s) // [a b c d e f g]

	s = slices.Delete(s, 0, 3)
	fmt.Println(s) // [d e f g]

	users := []user{{"Tami", 25}, {"Raul", 30}, {"Tony", 35}}

	users = slices.DeleteFunc(users, func(u user) bool {
		return u.age <= 30
	})
	fmt.Println(users) // [{Tony 35}]
}
