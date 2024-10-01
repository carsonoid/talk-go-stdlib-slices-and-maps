package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for k, v := range maps.All(m) {
		fmt.Println(k, v)
	}

	fmt.Println("Iterating over keys:")
	for k := range maps.Keys(m) {
		fmt.Println(k)
	}

	fmt.Println("Iterating over values:")
	for v := range maps.Values(m) {
		fmt.Println(v)
	}
}
