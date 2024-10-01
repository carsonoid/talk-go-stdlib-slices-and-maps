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
	fmt.Println(slices.Index(letters, "d")) // 3
	fmt.Println(slices.Index(letters, "z")) // -1

	users := []user{{"Tami", 25}, {"Raul", 30}, {"Tony", 35}}
	fmt.Println(slices.IndexFunc(users, findUser("Raul"))) // 1
}

func findUser(name string) func(user) bool {
	return func(u user) bool {
		return u.name == name
	}
}
