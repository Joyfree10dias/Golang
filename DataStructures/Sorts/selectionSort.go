package main

import(
	"fmt"
)

func selectionSort(l *[]int, length int) {
	list := *l
	for i := 0 ; i < length - 1 ; i++ {
		for j := i + 1 ; j < length ; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
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
	selectionSort(&numberList, len(numberList))
}