package main

import (
	"cmp"
	"fmt"
)

func main() {
	fmt.Println(
		cmp.Compare(0, 1), // -1
		cmp.Compare(0, 0), // 0
		cmp.Compare(1, 0), // 1
		cmp.Less(0, 1),    // true
		cmp.Less(0, 0),    // false
		cmp.Less(1, 0),    // false
	)

	fmt.Println(
		cmp.Compare("a", "b"), // -1
		cmp.Compare("a", "a"), // 0
		cmp.Compare("b", "a"), // 1
		cmp.Less("a", "b"),    // true
		cmp.Less("a", "a"),    // false
		cmp.Less("b", "a"),    // false
	)
}
