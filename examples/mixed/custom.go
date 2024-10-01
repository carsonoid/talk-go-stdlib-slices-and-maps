package main

import (
	"fmt"
	"iter"
	"slices"
)

// START OMIT
type user struct {
	name  string
	age   int
	roles []string
}

func (u *user) Roles() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, role := range u.roles {
			if !yield(role) {
				break
			}
		}
	}
}

// END OMIT

func main() {
	users := []user{
		{"Tami", 25, []string{"admin", "editor"}},
		{"Raul", 30, []string{"editor"}},
		{"Tony", 35, []string{"admin"}},
	}

	var allRoles []string
	for _, u := range users {
		allRoles = slices.AppendSeq(allRoles, u.Roles())
	}
	fmt.Println(allRoles) // [admin editor editor admin]

	clean := slices.Compact(slices.Sorted(slices.Values(allRoles)))
	fmt.Println(clean) // [admin editor]
}
