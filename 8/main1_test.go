package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestWaitGroup_Basic(t *testing.T) {
	wg := NewWaitGroup()

	var completed atomic.Int64
	n := 5

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond)
			completed.Add(1)
		}()
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		if completed.Load() != int64(n) {
			t.Errorf("завершилось %d горутин, ожидалось %d", completed.Load(), n)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("таймаут: Wait() не вернулся за 5 секунд")
	}
}

func TestWaitGroup_ZeroAdd(t *testing.T) {
	wg := NewWaitGroup()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Wait() заблокировался при нулевом счётчике")
	}
}
