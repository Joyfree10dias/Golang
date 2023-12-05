package main

import(
	"fmt"
)

// Declare Queue
type queue[T comparable] struct {
	Head, tail *element[T]
}

// Declare element 
type element[T comparable] struct {
	Data T
	next *element[T]
}

// Appends or attaches new elements at the end of the Queue
func (q *queue[T])enqueue(val T) {
	fmt.Println("Appended :", val)
	if q.Head == nil {
		q.Head = &element[T]{ Data: val }
		q.tail = q.Head
		q.Head.next = nil
		q.tail.next = nil
	} else {
		q.tail.next = &element[T]{ Data: val }
		q.tail = q.tail.next
		q.tail.next = nil
	}
}

func (q *queue[T])dequeue() {
	if q.Head == nil {
		fmt.Println("Dequeue failed!! No more elements in the queue.")
		return
	} else {
		fmt.Println("Removed :",q.Head.Data)
		q.Head = q.Head.next
	} 
}

// Displays the Queue elements
func (q *queue[T])Display() {
	curr := q.Head
	fmt.Printf("\nQueue : ")
	for curr != nil {
		fmt.Printf("%v ---> ", curr.Data)
		curr = curr.next
	}
	fmt.Printf("nil\n\n")
}

func main() {
	// Initialise a Queue 
	var Queue queue[int]
	Queue.enqueue(1)
	Queue.enqueue(2)
	Queue.enqueue(3)
	Queue.enqueue(6)
	Queue.Display()
	Queue.dequeue()
	Queue.dequeue()
	Queue.Display()	
	Queue.dequeue()
	Queue.dequeue()
	Queue.dequeue()
	Queue.Display()	
}