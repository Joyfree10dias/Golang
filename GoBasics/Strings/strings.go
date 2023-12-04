package main

import (
	"fmt"
	"strconv"
)

// This function checks whether a string has a charcter 'c'
func checkc(r rune) {
	if r == 'c' {
		fmt.Println("Found c")
	}
}


func main() {
	// A Go string is a read-only slice of bytes
	// The concept of a character is called a rune - it’s an integer that represents a Unicode code point
	const S = "A/สวัสดีc"
	// A string in Go is a bunch of bytes 
	
	// Using a normal for loop - prints the bytes
	for i := 0; i < len(S); i++{
		fmt.Printf("char : %q\n", S[i])
	}

	fmt.Println()
	// Use for range loop - to access rune
	for idx, runeValue := range S {
		fmt.Printf("rune : %s at position : %v\n", strconv.QuoteRune(runeValue), idx)
		checkc(runeValue)
	}

}