package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	// reverse
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	fmt.Println(a)

	// Insert n elements at position i:
	n := 3
	i := 2
	a = append(a[:i], append(make([]int, n), a[i:]...)...)
	fmt.Println(a)

	// Insert x at position i:
	x := 9
	a = append(a[:i], append([]int{x}, a[i:]...)...)
	fmt.Println(a)
}
