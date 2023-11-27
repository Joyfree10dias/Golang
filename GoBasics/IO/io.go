package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var firstName string
	// For a single value 
	fmt.Print("Enter first name: ")
	// Take in user input 
	fmt.Scan(&firstName)
	fmt.Printf("User name: %s\n",firstName)

	// For a line 
	// Method 1
	fmt.Print("Enter name (using NewScanner): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstName = scanner.Text()
	fmt.Printf("User name: %s\n",firstName)

	// Method 2 - return '\n' as well with the input
	fmt.Print("Enter name(using NewReader): ")
	reader := bufio.NewReader(os.Stdin)
	firstName, _ = reader.ReadString('\n')
	fmt.Printf("User name: %s",firstName)

}