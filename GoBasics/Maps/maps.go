package main

import (
	"fmt"
	"maps"
)

func main(){
	// Declare map - this map conatains an item and its price
	// Method 1
	var Items = map[string]int{"Lays": 10, "Parle-G": 10}
	// in this map, key --> item and value --> price 
	// add elements
	Items["Eggs"] = 60
	Items["Pop Corn"] = 30
	fmt.Println("Items :",Items)

	// Method 2
	cars := make(map[string]int)
	cars["BMW"] = 1000000
	cars["Alto"] = 100000
	fmt.Println("Cars :",cars)

	// Printing one item price
	itemPrice := Items["Lays"]
	fmt.Printf("Price of Lays is Rs. %v\n", itemPrice)

	// Print length of the map
	fmt.Println("No. of items :",len(Items))

	// Delete one item - use 'delete' keyword
	delete(Items, "Pop Corn")
	fmt.Println("Items :",Items)

	// Clear the Items map - use 'clear' keyword
	clear(Items)
	fmt.Println("Items :",Items)

	// Comparing two maps
	Map1 := map[int]int{0: 10, 1: 20}
	Map2 := map[int]int{0: 10, 1: 20}
	
	fmt.Println("Are maps equal?", maps.Equal(Map1, Map2))

	// Declare a list of maps
	listOfMaps := make([]map[int]int, 0)
	something := make(map[int]int)
	something[0] = 1
	something[1] = 2
	fmt.Println("Something :",something)

	listOfMaps = append(listOfMaps, something)
	listOfMaps = append(listOfMaps, something)
	listOfMaps = append(listOfMaps, something)
	fmt.Println("List of maps :",listOfMaps)
}