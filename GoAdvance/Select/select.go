package main

import (
	"fmt"
	"time"
)


// Go’s select lets you wait on multiple channel operations.

func main() {
	// Make channels 
	c1 := make(chan string)
	c2 := make(chan string)

	// Using closure 
	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0 ; i < 2 ; i++ {
		select{
		case msg := <- c1:
			fmt.Println("Message :",msg)
		case msg := <- c2:
			fmt.Println("Message :",msg)
		}
	}
}