package main

import (
	"testing"
	"time"
)

func TestRandomIntGenerator(t *testing.T) {
	rng := RandomIntGenerator()

	for i := 0; i < 100; i++ {
		select {
		case val := <-rng:
			if val < 0 || val >= 1000 {
				t.Errorf("значение %d вне диапазона [0, 1000)", val)
			}
		case <-time.After(time.Second):
			t.Fatal("таймаут: канал не вернул значение за 1 секунду")
		}
	}
}
