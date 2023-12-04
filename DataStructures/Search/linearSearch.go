package main

import (
	"fmt"
	"strconv"
)

// Linear search we check if the element we want is there in the list by comparing it with every number in it
func linearSearch(ptr *[]int, length int, element int) string {
	Pointer := *ptr
	for i := 0 ; i < length ; i++ {
		if Pointer[i] == element {
			return strconv.FormatInt(int64(element), 10) + " found at position " + strconv.FormatInt(int64(i + 1), 10)
		}
	}
	return strconv.FormatInt(int64(element), 10) + " not found"
}

func main() {
	// Declare slice 
	var numberList = []int{ 10, 3, 6, 78, 34, 89, 1, 0, 45 }
	fmt.Println("Number List :", numberList)

	// Call linearSearch 
	fmt.Println(linearSearch(&numberList, len(numberList), 10))
}