package main

import (
	"errors"
	"slices"
	"testing"
)

func TestNewOriginalSliceResult(t *testing.T) {
	t.Run("new original slice", func(t *testing.T) {
		slice := newOriginalSlice()
		needleSliceLen := 10

		if len(slice) != needleSliceLen {
			t.Errorf("newOriginalSlice: len:%d needle:%d", len(slice), needleSliceLen)
		}
	})
}

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"has even", []int{1, 2, 3, 4}, []int{2, 4}},
		{"doesnt has even", []int{1, 3, 5}, []int{}},
		{"empty slice", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceExample(tt.input)

			if !slices.Equal(got, tt.want) {
				t.Errorf("sliceExample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		add   int
		want  []int
	}{
		{"add element 11", []int{1, 2}, 11, []int{1, 2, 11}},
		{"add zero", []int{1, 2}, 0, []int{1, 2, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addElements(tt.input, tt.add)

			if !slices.Equal(got, tt.want) {
				t.Errorf("addElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	t.Run("copy slice", func(t *testing.T) {
		got := copySlice([]int{1, 2, 3})

		if !slices.Equal(got, []int{1, 2, 3}) {
			t.Errorf("copySlice() = %v, want %v", got, []int{1, 2, 3})
		}
	})
}

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		index int
		want  []int
		err   error
	}{
		{"remove index 1", []int{1, 2, 3, 4}, 1, []int{1, 3, 4}, nil},
		{"remove index 0", []int{1, 2}, 0, []int{2}, nil},
		{"out of range index", []int{1, 2}, 3, nil, errors.New("error: index out of range")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeElement(tt.input, tt.index)

			if err != nil && !errors.Is(err, ErrIndexOutOfRange) {
				t.Errorf("got %v, want %v", err, ErrIndexOutOfRange)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
