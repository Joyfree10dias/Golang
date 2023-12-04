package main

import(
	"fmt"
	"strconv"
)

// To perform binary search the list has to be sorted
func binarySearch(list []int, length, element int) string {
	mid := 0
	low := 0
	high := length - 1 
	for low <= high {
		mid = (low + high) / 2
		if element > list[mid] {
			low = mid + 1
		} else if element == list[mid] {
			return strconv.FormatInt(int64(element), 10) + " found at position " + strconv.FormatInt(int64(mid + 1), 10)
		} else {
			high = mid - 1  
		}
	}

	return strconv.FormatInt(int64(element), 10) + " not found"
}

func main() {
	var numberList = []int{ 1, 2, 4, 6, 8, 12, 45, 67}
	fmt.Println("Number List :",numberList)

	// Call binarySearch 
	fmt.Println(binarySearch(numberList, len(numberList), 10))
}