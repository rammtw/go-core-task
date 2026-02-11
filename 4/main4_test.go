package main

import (
	"slices"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		s1       []string
		s2       []string
		expected []string
	}{
		{
			"slice diff",
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{"banana", "date", "fig"},
			[]string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			"all diff",
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{},
		},
		{
			"not contains elements",
			[]string{"apple", "banana"},
			[]string{"lead", "gno1"},
			[]string{"apple", "banana"},
		},
		{
			"empty",
			[]string{},
			[]string{},
			[]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Difference(test.s1, test.s2)

			if !slices.Equal(got, test.expected) {
				t.Errorf("Difference() = %v, want %v", got, test.expected)
			}
		})
	}
}
