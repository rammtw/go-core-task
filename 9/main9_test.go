package main

import (
	"math"
	"testing"
	"time"
)

func TestGenerate_ProducesValues(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	out := generate(done)

	for i := 0; i < 20; i++ {
		select {
		case v := <-out:
			if v > 255 {
				t.Errorf("значение %d выходит за диапазон uint8", v)
			}
		case <-time.After(time.Second):
			t.Fatal("таймаут: generate не отдал значение за 1 секунду")
		}
	}
}

func TestCube_CorrectValues(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	in := make(chan uint8)
	out := cube(done, in)

	tests := []struct {
		input    uint8
		expected float64
	}{
		{0, 0},
		{1, 1},
		{2, 8},
		{3, 27},
		{5, 125},
		{10, 1000},
		{255, 16581375},
	}

	for _, tc := range tests {
		in <- tc.input
		select {
		case result := <-out:
			if math.Abs(result-tc.expected) > 0.001 {
				t.Errorf("cube(%d) = %.0f, ожидалось %.0f", tc.input, result, tc.expected)
			}
		case <-time.After(time.Second):
			t.Fatalf("таймаут при обработке %d", tc.input)
		}
	}
}
