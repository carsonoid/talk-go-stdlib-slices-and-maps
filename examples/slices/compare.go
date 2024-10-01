package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
)

func main() {
	letters := []string{"a", "b", "c"}
	strings := []string{"1", "2", "3"}
	ints := []int{1, 2, 3}

	fmt.Println(slices.Compare(letters, strings)) // 1 because "a" > "1"

	fn := func(s string, i int) int {
		n, err := strconv.Atoi(s)
		if err != nil {
			return 1
		}
		return cmp.Compare(n, i)
	}
	fmt.Println(slices.CompareFunc(strings, ints, fn)) // 0 because 1 == 1, 2 == 2, 3 == 3
}
