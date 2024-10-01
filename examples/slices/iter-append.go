package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"f", "g", "e"}

	s3 := slices.AppendSeq(s1, slices.Values(s2))
	fmt.Println(s3)
}
