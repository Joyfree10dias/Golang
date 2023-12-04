package main

import (
	"fmt"
)

// Generics 
// This function takes a map as input(Generic) and returns a slice which contains map keys  
func mapKeys[K comparable, V any] (m map[K]V) []K {
	var keySlice = []K{}
	for key := range m {
		keySlice = append(keySlice, key)
	}
	return keySlice
}

// This struct stores some data and pointer to that struct 
type myData[T any] struct{
	Data T
	Pointer *myData[T]
	integer int
}

// This function adds two variables
func add[T int | string | float64](a T, b T) T {
	return a + b
}


func main() {
	// Get the keys of the given map  
	var m = map[string]string{ "3": "Dog", "2": "Cat", "0": "Donkey" }
	res := mapKeys(m)
	fmt.Println("Result :",res)

	// T is integer 
	var D = myData[int]{ Data: 123, integer: 1234}
	D.Pointer = &D
	fmt.Println("My Data :",D)

	// T is string 
	D1 := myData[string]{ Data: "123", integer: 1234}
	D1.Pointer = &D1
	fmt.Println("My Data :",D1)

	// Adding two integers
	resultInt := add(1, 2)  
	fmt.Println("Result :",resultInt)

	// Adding two strings
	var resultStr = add("1", "2")  
	fmt.Println("Result :",resultStr)

	// Adding two strings
	var resultFloat = add(1.3, 2.2)  
	fmt.Println("Result :",resultFloat)
}