package main

import(
	"fmt"
)

// Bubble Sort 
func bubbleSort(l *[]int, length int) {
	list := *l
	for i := 0 ; i < length - 1 ; i++ {
		for j := 0 ; j < length - i - 1 ; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
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
	// Call Bubble Sort 
	bubbleSort(&numberList, len(numberList))
}