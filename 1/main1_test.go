package main

import (
	"testing"
)

func TestBuildStringResult(t *testing.T) {
	tests := []struct {
		name           string
		input          Params
		expectedString string
	}{
		{"Task 1", Params{
			IntDec:  42,
			IntOct:  052,
			IntHex:  0x2A,
			Float:   3.14,
			Str:     "Golang",
			Bool:    true,
			Complex: complex64(1 + 2i),
		}, "4242423.14000Golangtrue(1.00+2.00i)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildString(tt.input)
			if result != tt.expectedString {
				t.Errorf("BuildString\n got: %q \nwant: %q", result, tt.expectedString)
			}
		})
	}
}

func TestInsertSalt(t *testing.T) {
	tests := []struct {
		name string
		str  string
		salt string
		want string
	}{
		{"Even", "qwer", "AA", "qwAAer"},
		{"Odd", "qwert", "SS", "qwSSert"},
		{"Empty salt", "qwerty", "", "qwerty"},
		{"Empty str", "", "salt", "salt"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InsertSalt([]rune(tt.str), tt.salt)
			if string(result) != tt.want {
				t.Errorf("InsertSalt\n got: %q \nwant: %q", result, tt.want)
			}
		})
	}
}

func TestEncodeStringSha256(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			"Empty",
			"",
			"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			"Test 1",
			"hello world",
			"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EncodeStringSha256(tt.str)
			if result != tt.want {
				t.Errorf("EncodeStringSha256\n got: %q \nwant: %q", result, tt.want)
			}
		})
	}
}
