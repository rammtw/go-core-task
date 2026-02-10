package main

import (
	"maps"
	"reflect"
	"testing"
)

func TestNewStringIntMap(t *testing.T) {
	t.Run("create new string int map", func(t *testing.T) {
		got := NewStringIntMap()

		reflectType := reflect.TypeOf(got)

		if reflectType != reflect.TypeOf(StringIntMap{}) {
			t.Errorf("NewStringIntMap() did not return a new string int map")
		}
	})
}

func TestStringIntMap_Add(t *testing.T) {
	tests := []struct {
		name string
		m    StringIntMap
		k    string
		v    int
		want StringIntMap
	}{
		{"add value", StringIntMap{}, "one", 1, StringIntMap{"one": 1}},
		{"add value2", StringIntMap{"c": 0, "a": 1}, "x", 0, StringIntMap{"a": 1, "c": 0, "x": 0}},
		{"add empty key", StringIntMap{"c": 0, "a": 1}, "", 0, StringIntMap{"a": 1, "c": 0, "": 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Add(tt.k, tt.v)

			if !maps.Equal(tt.m, tt.want) {
				t.Errorf("Add() = %v, want %v", tt.m, tt.want)
			}
		})
	}
}

func TestStringIntMap_Remove(t *testing.T) {
	tests := []struct {
		name string
		m    StringIntMap
		k    string
		want StringIntMap
	}{
		{"remove undefined index", StringIntMap{}, "hash", StringIntMap{}},
		{"remove index", StringIntMap{"first": 0}, "first", StringIntMap{}},
		{"remove index 2", StringIntMap{"elem": 22, "elem_two": 33, "zero": 100}, "elem", StringIntMap{"elem_two": 33, "zero": 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Remove(tt.k)

			if !maps.Equal(tt.m, tt.want) {
				t.Errorf("Remove() = %v, want %v", tt.m, tt.want)
			}
		})
	}
}

func TestStringIntMap_Copy(t *testing.T) {
	t.Run("copy", func(t *testing.T) {
		m := NewStringIntMap()
		s := m.Copy()
		m.Add("z", 1)
		s.Add("x", 2)

		if _, ok := m["x"]; ok {
			t.Errorf("NewStringIntMap() did not copy")
		}
		if _, ok := s["z"]; ok {
			t.Errorf("NewStringIntMap() did not copy")
		}
	})
}

func TestStringIntMap_Exists(t *testing.T) {
	tests := []struct {
		name  string
		m     StringIntMap
		index string
		want  bool
	}{
		{"exists", StringIntMap{"key": 1}, "key", true},
		{"not exists", StringIntMap{"one": 2}, "key", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Exists(tt.index); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStringIntMap_Get(t *testing.T) {
	tests := []struct {
		name      string
		m         StringIntMap
		index     string
		wantValue int
		ok        bool
	}{
		{"get exists", StringIntMap{"key": 1}, "key", 1, true},
		{"get not exists", StringIntMap{"key": 1}, "key2", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOk := tt.m.Get(tt.index)
			if gotValue != tt.wantValue || gotOk != tt.ok {
				t.Errorf("Get() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
