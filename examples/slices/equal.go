package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	letters := []string{"a", "b", "c"}
	stringNums := []string{"1", "2", "3"}
	intNums := []int{1, 2, 3}

	fmt.Println(slices.Equal(letters, stringNums)) // false

	fn := func(s string, i int) bool {
		n, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return n == i
	}
	fmt.Println(slices.EqualFunc(stringNums, intNums, fn)) // true
}
