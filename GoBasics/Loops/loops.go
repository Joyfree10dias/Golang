package main

import (
	"fmt"
)

func main() {
	// In 'GO' we have only one kind of loop which can be used for different cases 
	// Infinite Loop
	// for {
	// 	fmt.Println("Infinite")
	// }

	// While loop - Printing 5 Numbers 
	n := 5
	i := 1
	for i <= n {
		fmt.Println(i)
		i++
	}

	// For Each 
	var Items = []string {"Bat", "Ball", "Pen" , "Book", "Shirt"}
	for Index, Item := range Items {
		fmt.Printf("Found item : %v at poistion %v\n", Item, Index)
	}

	for i2, j := 0, len(Items)-1; i2 < j; i2, j = i2+1, j-1 {
		Items[i2], Items[j] = Items[j], Items[i2]
	}

	// For loop 
	for idx, j := 0, len(Items) - 1; idx < j; idx, j = idx + 1, j - 1{
		fmt.Printf("i : %v and j : %v\n", idx, j)
	}

}