package main

import (
	"fmt"
)

func main() {
	var city string = "San Fransisco"
	// Define a switch case - to identify which city belongs to which country
	switch city{
	case "Mumbai", "Chennai", "Delhi":
		fmt.Println("Country : India")
	case "New York", "San Fransisco":
		fmt.Println("Country : USA")
	case "London", "Manchester":
		fmt.Println("Country : UK")
	case "Toronto":
		fmt.Println("Country : Canada")
	}

}