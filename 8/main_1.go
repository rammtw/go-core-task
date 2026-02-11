package main

import (
	"fmt"
	"sync/atomic"
)

type WaitGroup struct {
	counter atomic.Int64
	done    chan struct{}
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		done: make(chan struct{}),
	}
}

func main() {
	wg := NewWaitGroup()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println(i)
		}()
	}

	wg.Wait()
}

func (wg *WaitGroup) Add(delta int) {
	wg.counter.Add(int64(delta))
}

func (wg *WaitGroup) Done() {
	if wg.counter.Add(-1) == 0 {
		close(wg.done)
	}
}

func (wg *WaitGroup) Wait() {
	if wg.counter.Load() == 0 {
		return
	}
	<-wg.done
}
