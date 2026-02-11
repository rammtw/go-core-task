package main

import (
	"testing"
	"time"
)

func TestMergeChannels_TotalCount(t *testing.T) {
	ch1 := producer("Chan#1", 15)
	ch2 := producer("Chan#2", 30)
	ch3 := producer("Chan#3", 45)

	merged := MergeChannels(ch1, ch2, ch3)
	var totalCount int
	timeout := time.After(5 * time.Second)

	done := false
	for !done {
		select {
		case _, ok := <-merged:
			if !ok {
				done = true
			} else {
				totalCount++
			}
		case <-timeout:
			t.Fatal("тест завис: канал не закрылся за 5 секунд")
		}
	}

	expected := 15 + 30 + 45
	if totalCount != expected {
		t.Errorf("Merged channel count was incorrect, got: %d, want: %d.", totalCount, expected)
	}
}
