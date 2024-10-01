package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	maps.DeleteFunc(m, func(k string, v int) bool {
		return k == "b" || v == 3
	})
	fmt.Println(m) // map[a:1]
}
