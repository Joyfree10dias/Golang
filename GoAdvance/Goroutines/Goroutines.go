package main

import(
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution.

func f(from string) {
	for i := 0 ; i < 3 ; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Call f()
	f("direct")

	// Call f() using Goroutines(go)
	go f("Goroutine")

	// Using closure 
	go func(msg string) {
		fmt.Println("Message :", msg)
	}("Going")

	// Sleep for 1s 
	time.Sleep(time.Second)
	fmt.Println("Done")
}