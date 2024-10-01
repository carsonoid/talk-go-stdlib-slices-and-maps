package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	ints := []int{1, 2, 3}
	purchases := []purchase{
		{2, "second"},
		{3, "first"},
		{1, "third"},
	}

	fmt.Println(slices.Min(ints))
	fmt.Println(slices.MinFunc(purchases, cmpPurchases))
	fmt.Println(slices.Max(ints))
	fmt.Println(slices.MaxFunc(purchases, cmpPurchases))
}

type purchase struct {
	priority int
	name     string
}

func cmpPurchases(a, b purchase) int {
	return cmp.Compare(a.priority, b.priority)
}
