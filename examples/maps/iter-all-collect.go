package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	for k, v := range maps.All(m) {
		fmt.Println(k, v)
	}

	maps.Insert(m, maps.All(map[string]int{"f": 6, "g": 7}))

	m2 := maps.Collect(maps.All(m))
	fmt.Println(m2) // map[a:1 b:2 c:3 d:4 e:5]
}
