package main

import (
	"fmt"
	"slices"
)

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(Difference(slice1, slice2))
}

func Difference(s, rem []string) []string {
	out := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if !slices.Contains(rem, s[i]) {
			out = append(out, s[i])
		}
	}
	return out
}
