package main

import (
	"fmt"
)

func main() {

	// Conditions
	// <, >, ==, >=, <=, &&, ||
	var choice int = 1
	if choice == 0 {
		fmt.Println("You entered choice :",choice)
	} else if choice == 1 {
		fmt.Println("You entered choice :",choice)
	} else {
		fmt.Println("Your choice is invalid")
	}

	// using OR - only one condition has to be true if we want enter the block 
	if choice < 2 || choice == 0 {
		fmt.Println("Your choice is valid")
	}

	// using AND - both conditions have to be true if we want enter the block 
	if choice < 2 && choice == 1 {
		fmt.Println("You entered choice :",choice)
	}
}