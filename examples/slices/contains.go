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
	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slices.Contains(letters, "b")) // true

	users := []user{{"Tami", 25}, {"Raul", 30}, {"Tony", 35}}

	fmt.Println(slices.ContainsFunc(users, func(u user) bool {
		return u.name == "Raul"
	})) // true

	fn := findUser("Raul")
	fmt.Println(slices.ContainsFunc(users, fn)) // true
}

func findUser(name string) func(user) bool {
	return func(u user) bool {
		return u.name == name
	}
}
