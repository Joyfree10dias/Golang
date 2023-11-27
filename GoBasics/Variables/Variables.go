package main

import (
	"fmt"
)

func main(){
	// 1st approach
	var name string = "Joyfree Dias"
	// 2nd approach - syntactic sugar
	age := 56
	fmt.Println("User name :", name, "and age :",age)

	// Constants
	const gender = "male"
	fmt.Println("Gender :",gender)

	// Using Printf 
	// %v (default) - Placeholder 
	fmt.Printf("Using Printf (default): \nUser name : %v and age : %v\n", name, age)
	// %s - string and %d - integer 
	fmt.Printf("Using Printf: \nUser name : %s and age : %d\n", name, age)

	var message string  
	message = "I love coding\n"
	fmt.Printf("Message : %v",message)
}