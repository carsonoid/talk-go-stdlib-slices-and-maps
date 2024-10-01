package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	purchases := []purchase{
		{3, "first"},
		{2, "second"},
		{1, "third"},
		{2, "fourth"},
	}

	fmt.Println("Sorted      ", slices.SortedFunc(slices.Values(purchases), sortOrder))
	fmt.Println("SortedStable", slices.SortedStableFunc(slices.Values(purchases), sortOrder))
}

// START OMIT
type purchase struct {
	priority int
	name     string
}

func sortOrder(a, b purchase) int {
	return cmp.Compare(a.priority, b.priority)
}

// END OMIT
