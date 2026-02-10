package main

import (
	"fmt"
)

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println(HasIntersect(a, b))
}

func HasIntersect(a, b []int) (bool, []int) {
	exclude := make(map[int]struct{}, len(b))
	for _, v := range b {
		exclude[v] = struct{}{}
	}
	intersect := make([]int, 0)
	for _, v := range a {
		if _, found := exclude[v]; found {
			intersect = append(intersect, v)
		}
	}
	if len(intersect) == 0 {
		return false, intersect
	}
	return true, intersect
}
