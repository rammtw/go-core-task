package main

import (
	"fmt"
	"maps"
)

type StringIntMap map[string]int

func main() {
	m := NewStringIntMap()

	m.Add("how", 2)
	m.Add("hello", 4)
	m.Remove("how")

	s := m.Copy()

	m.Add("New to Original", 100)
	s.Add("New to Copy", 100)

	ok := s.Exists("hello")
	fmt.Println("Original:", m)
	fmt.Println("Copy", s)
	fmt.Println("Has OK key in copy: ", ok)
}

func NewStringIntMap() StringIntMap {
	return make(StringIntMap)
}

func (m StringIntMap) Add(key string, value int) {
	m[key] = value
}

func (m StringIntMap) Remove(key string) {
	delete(m, key)
}

func (m StringIntMap) Copy() StringIntMap {
	newMap := make(StringIntMap)
	maps.Copy(newMap, m)
	return newMap
}

func (m StringIntMap) Exists(key string) bool {
	_, ok := m[key]
	return ok
}

func (m StringIntMap) Get(key string) (int, bool) {
	if value, ok := m[key]; ok {
		return value, ok
	}
	return 0, false
}
