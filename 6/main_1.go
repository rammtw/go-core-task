package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rng := RandomIntGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(<-rng)
	}
}

func RandomIntGenerator() <-chan int {
	c := make(chan int)
	go func() {
		for {
			c <- rand.Intn(1000)
		}
	}()
	return c
}
