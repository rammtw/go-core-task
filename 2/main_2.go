package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrIndexOutOfRange = errors.New("error: index out of range")

func main() {
	originalSlice := newOriginalSlice()

	s := sliceExample(originalSlice)
	fmt.Println("Только четные:", s)

	s = addElements(s, 11)
	s = addElements(s, 12)

	fmt.Println("Добавлены новые элементы:", s)

	v := copySlice(s)

	fmt.Printf("Копия слайса: %v, addr1: %p, addr2: %p\n", v, s, v)

	removeIndex := rand.Intn(len(v))
	d, err := removeElement(v, removeIndex)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("После удаления %d индекса: %v\n", removeIndex, d)
	fmt.Println("Исходный слайс:", originalSlice)
}

func newOriginalSlice() []int {
	s := make([]int, 0)
	for i := 0; i < 10; i++ {
		s = append(s, rand.Intn(50))
	}
	return s
}

func sliceExample(s []int) []int {
	out := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i]%2 == 0 {
			out = append(out, s[i])
		}
	}
	return out
}

func addElements(s []int, n int) []int {
	return append(s, n)
}

func copySlice(s []int) []int {
	newSlice := make([]int, len(s))
	copy(newSlice, s)
	return newSlice
}

func removeElement(s []int, i int) ([]int, error) {
	if i < 0 || i >= len(s) {
		return nil, ErrIndexOutOfRange
	}

	newSlice := make([]int, 0)

	newSlice = append(newSlice, s[:i]...)
	newSlice = append(newSlice, s[i+1:]...)

	return newSlice, nil
}
