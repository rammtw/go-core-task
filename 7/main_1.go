package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := producer("Chan#1", 5)
	ch2 := producer("Chan#2", 2)
	ch3 := producer("Chan#3", 4)
	ch4 := producer("Chan#4", 6)

	merged := MergeChannels(ch1, ch2, ch3, ch4)

	for msg := range merged {
		fmt.Println(msg)
	}
}

func MergeChannels(channels ...<-chan string) chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, c := range channels {
		go func() {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func producer(name string, count int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			ch <- fmt.Sprintf("%s: сообщение %d", name, i)
		}
	}()
	return ch
}
