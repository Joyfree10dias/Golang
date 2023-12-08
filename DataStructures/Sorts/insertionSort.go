package main

import (
	"fmt"
)

// Insertion Sort
func insertionSort(l *[]int, length int) {
	list := *l
	for i := 0 ; i < length ; i++ {
		for j := 0 ; j < i ; j++ {
			if list[j] > list[i] {
				var temp int = 0
				temp = list[i]
				for idx := i ; j < idx ; idx-- {
					list[idx] = list[idx - 1]
				}
				list[j] = temp
				break
			}
		}
	}
	// To display list 
	fmt.Printf("Display : ")
	for i := 0 ; i < length ; i++ {
		fmt.Printf("%d ",list[i])
	}
	fmt.Println()
}

func main() {
	// Initialise list 
	var numberList = []int {3, 5, 2, 78, 45, 1, 4, 30}
	fmt.Println("Number List :", numberList)
	insertionSort(&numberList, len(numberList))
}