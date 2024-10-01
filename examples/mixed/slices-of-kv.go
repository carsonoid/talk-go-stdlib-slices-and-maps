package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	keys := slices.Collect(maps.Keys(m))
	fmt.Println(keys) // [a b c d e]

	values := slices.Collect(maps.Values(m))
	fmt.Println(values) // [1 2 3 4 5]
}
