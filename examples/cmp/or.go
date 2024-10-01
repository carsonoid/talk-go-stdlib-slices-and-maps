package main

import (
	"cmp"
	"fmt"
	"os"
)

var addFlag = ""

func main() {
	fmt.Println(cmp.Or("", "a", "b")) // a
	fmt.Println(cmp.Or(0, 0, 1, 2))   // 1

	// practical use case: config
	addr := cmp.Or(os.Getenv("ADDR"), addFlag, "localhost")
	fmt.Println(addr)
}
