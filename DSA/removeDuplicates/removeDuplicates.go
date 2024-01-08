package main

import (
	"fmt"
	"time"
)

// Removes Duplicate elemenets from the sorted array and returns the no. of unique elements
func removeDuplicates(nums []int) (int, []int) {
	// Approach 1 :
    nextInsertionIndex := 1
    
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            nums[nextInsertionIndex] = nums[i]
            nextInsertionIndex++
        }
    }
    return nextInsertionIndex, nums[:nextInsertionIndex]

	// Approach 2 :
	// i := 0
    // for j := range nums {
    //     if nums[i] != nums[j] {
    //         i += 1
    //         nums[i] = nums[j]
    //     }
    // }
    // return i + 1, nums[:i + 1]
}

func main() {
	// Define nums (sorted) 
	// var nums = []int{1, 1, 3, 4, 5, 5, 6, 7, 7, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	// var nums = []int{1, 1, 1, 1, 3, 4, 5, 5, 6, 7, 7, 7, 7}

	// Call removeDuplicate
	start := time.Now()
	for i := 0 ; i < 100000 ; i++ { 
		var nums = []int{1, 1, 3, 4, 5, 5, 6, 7, 7, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
		removeDuplicates(nums)
		// fmt.Println("Output :", k, ",", array)
	}
	
	end := time.Now()
	dur := end.Sub(start) 
	fmt.Println("Time :", dur.Seconds(), "s")
}