package main

import (
	"fmt"
	"maps"
	"strconv"
)

type user struct {
	name string
	age  int
}

func main() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println(maps.Equal(m1, m2)) // true

	m3 := map[string]string{"a": "1", "b": "2", "c": "3"}
	fmt.Println(maps.EqualFunc(m1, m3, cmpIntStr)) // true

	users1 := map[string]user{"a": {"Tami", 25}, "b": {"Raul", 30}}
	users2 := map[string]user{"b": {"Raul", 11}, "a": {"Tami", 22}}
	fmt.Println(maps.EqualFunc(users1, users2, cmpUser)) // true
}

func cmpIntStr(v1 int, v2 string) bool {
	return strconv.Itoa(v1) == v2
}

func cmpUser(u1 user, u2 user) bool {
	return u1.name == u2.name
}
