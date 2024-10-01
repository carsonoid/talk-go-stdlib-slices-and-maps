package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	unsorted := []string{"b", "c", "a", "e", "d"}

	fmt.Println("Sorted          ", slices.Sorted(slices.Values(unsorted)))
	fmt.Println("SortedFunc      ", slices.SortedFunc(slices.Values(unsorted), cmp.Compare))
	fmt.Println("SortedStableFunc", slices.SortedStableFunc(slices.Values(unsorted), cmp.Compare))
}
