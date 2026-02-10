package main

import (
	"slices"
	"testing"
)

func TestHasIntersect(t *testing.T) {
	tests := []struct {
		name      string
		a         []int
		b         []int
		want      bool
		wantSlice []int
	}{
		{
			"first slice is longer", []int{1, 2, 3, 4}, []int{4, 5}, true, []int{4},
		},
		{
			"second slice is longer", []int{1, 2}, []int{4, 5, 2, 7}, true, []int{2},
		},
		{
			"no matches", []int{1, 2}, []int{4, 5}, false, []int{},
		},
		{
			"empty slices", []int{}, []int{}, false, []int{},
		},
		{
			"second empty", []int{1, 2}, []int{}, false, []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotSlice := HasIntersect(tt.a, tt.b)
			if gotOk != tt.want || !slices.Equal(gotSlice, tt.wantSlice) {
				t.Errorf("HasIntersect() = %v, want %v", gotSlice, tt.wantSlice)
			}
		})
	}
}
