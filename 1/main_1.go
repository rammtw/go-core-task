package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type IntOct int
type IntHex int

type Params struct {
	IntDec  int
	IntOct  IntOct
	IntHex  IntHex
	Float   float64
	Str     string
	Bool    bool
	Complex complex64
}

func main() {
	params := Params{
		IntDec:  42,
		IntOct:  0o52,
		IntHex:  0x2A,
		Float:   3.14159,
		Str:     "hello world",
		Bool:    true,
		Complex: complex64(3 + 4i),
	}

	fmt.Printf("intDec = %d,\t тип: %T\n", params.IntDec, params.IntDec)
	fmt.Printf("intOct = %d,\t тип: %T\n", params.IntOct, params.IntOct)
	fmt.Printf("intHex = %d,\t тип: %T\n", params.IntHex, params.IntHex)
	fmt.Printf("flt    = %g,\t тип: %T\n", params.Float, params.Float)
	fmt.Printf("str    = %s,\t тип: %T\n", params.Str, params.Str)
	fmt.Printf("flag   = %t,\t тип: %T\n", params.Bool, params.Bool)
	fmt.Printf("complexNum      = %v,\t тип: %T\n", params.Complex, params.Complex)

	result := BuildString(params)
	runes := StrToRun(result)
	salted := InsertSalt(runes, "go-2024")
	hash := EncodeStringSha256(string(salted))

	fmt.Printf("\nSHA256: %s\n", hash)
}

func EncodeStringSha256(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func InsertSalt(runes []rune, salt string) []rune {
	saltRunes := []rune(salt)
	mid := len(runes) / 2

	salted := make([]rune, 0, len(runes)+len(saltRunes))
	salted = append(salted, runes[:mid]...)
	salted = append(salted, saltRunes...)
	salted = append(salted, runes[mid:]...)

	return salted
}

func StrToRun(s string) []rune {
	return []rune(s)
}

func BuildString(p Params) string {
	parts := []string{
		strconv.Itoa(p.IntDec),
		strconv.Itoa(int(p.IntOct)),
		strconv.Itoa(int(p.IntHex)),
		strconv.FormatFloat(p.Float, 'f', 5, 64),
		p.Str,
		strconv.FormatBool(p.Bool),
		strconv.FormatComplex(complex128(p.Complex), 'f', 2, 64),
	}

	return strings.Join(parts, "")
}
