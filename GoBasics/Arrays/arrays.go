package main

import(
	"fmt"
)

func main(){
	// arrays in go have fixed size 
	// Declare an array
	// Method 1
	var family = [5]string {"Joyfree", "Joshan", "Joaquim"}
	fmt.Println("Family :",family)

	// Method 2
	var animals [5]string
	// Adding new elements 
	animals[0] = "Dog"
	animals[1] = "Cat"
	animals[3] = "Parrot"
	animals[2] = "Donkey"
	fmt.Println("Animals :",animals)

	// Slices(Dynamic version of arrays) - an abstraction of Arrays 
	var numbers []int
	// Adding new elements dynamically using append 
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	fmt.Println("Numbers :",numbers)

	// Removal of any element is similar to that of C or C++
	// Method 1
	// we need the index of the element to be removed 
	// after removal we need to shift all the other elemets towards left  
	// Method 2 
	numbers = removeElement(numbers, 2)
	fmt.Println("Numbers :",numbers)
}

// Removes element from the desired position
func removeElement(slice []int, position int) []int {
	return append(slice[:position-1], slice[position:]...)
}