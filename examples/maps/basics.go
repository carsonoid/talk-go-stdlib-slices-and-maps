package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m["d"] = 4     // adds a new key-value pair
	m["a"] = 0     // overwrites the previous value
	delete(m, "b") // deletes by key

	fmt.Println(m) // map[a:0 c:3 d:4]
}
