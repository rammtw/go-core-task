package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	done := make(chan struct{})

	numbers := generate(done)
	results := cube(done, numbers)

	for i := 0; i < 10; i++ {
		fmt.Printf("uint8: куб: %.0f\n", <-results)
	}

	close(done)
	fmt.Println("Конвейер остановлен")
}

func cube(done <-chan struct{}, in <-chan uint8) <-chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)
		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- math.Pow(float64(v), 3)
			case <-done:
				return
			}
		}
	}()
	return out
}

func generate(done <-chan struct{}) <-chan uint8 {
	out := make(chan uint8)
	go func() {
		defer close(out)
		for {
			select {
			case out <- uint8(rand.Intn(256)):
			case <-done:
				return
			}
		}
	}()
	return out
}
