package main

import (
	"fmt"
	"time"
)

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine
// and receive those values into another goroutine.

func main() {
	// Declare a channel
	messageChannel := make(chan string) // this is an unbuffered channel

	// Using closure
	go func() {
		// send a message in the message channel
		messageChannel <- "Ping"
		messageChannel <- "Me"
		fmt.Println("Sending Ping")
	}()

	// recieve a message from the message channel
	fmt.Println("Message channel :", messageChannel)
	message := <-messageChannel
	fmt.Println("Message :", message)
	message = <-messageChannel
	fmt.Println("Message :", message)

	// By default channels are unbuffered, meaning that they will only accept sends (chan <-)
	// if there is a corresponding receive (<- chan) ready to receive the sent value.
	// Buffered channels accept a limited number of values without a corresponding receiver for those values.

	// Buffered channel
	buffMessageChannel := make(chan string, 1)

	// Using closure
	go func() {
		buffMessageChannel <- "Something"
		buffMessageChannel <- "Nothing"
	}()

	syncChannel := make(chan bool, 1)
	// Call worker()
	go worker(syncChannel)
	
	fmt.Println("Address of syncChannel :",syncChannel)

	// If you removed the <- syncChannel line from this program, 
	// the program would exit before the worker even started.
	<- syncChannel

	// When using channels as function parameters, 
	// you can specify if a channel is meant to only send or receive values. 
	// This specificity increases the type-safety of the program.
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// Call ping() and pong()
	go ping(pings, "Passing message")
	go pong(pings, pongs)
	fmt.Println("Pongs :", <- pongs)
}

// This is the function we’ll run in a goroutine.
// The done channel will be used to notify another goroutine that this function’s work is done.
func worker(done chan bool) {
	fmt.Println("Address of done :",done)
	fmt.Printf("Working.... ")
	time.Sleep(time.Second * 3)
	fmt.Println("done")
	done <- true
}

// this functions only sends messages inside channel 
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Recieves the message and sends it to another channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <- pings
	pongs <- msg 
}

