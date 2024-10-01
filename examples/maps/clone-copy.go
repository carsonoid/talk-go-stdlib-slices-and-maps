package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := maps.Clone(m1)
	m1["c"] = 4
	fmt.Println(m1) // map[a:1 b:2 c:3]
	fmt.Println(m2) // map[a:1 b:2 c:4]

	m3 := map[string]int{"a": 11, "b": 12, "c": 13, "d": 14}
	maps.Copy(m1, m3)
	fmt.Println(m1) // map[a:11 b:12 c:13 d:14]
}
