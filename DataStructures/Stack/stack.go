package main

import(
	"fmt"
)

// Implementing stack using struct
// Declare stack struct 
type stack[T any] struct {
	Head, tail *element[T]
}

// Declare element
type element[T any] struct {
	Data T
	next *element[T]
}

// Pushes the element on to the stack
func (s *stack[T])push(val T) {
	fmt.Println("Pushed :",val)
	if s.Head == nil {
		s.Head = &element[T]{ Data: val }
		s.tail = s.Head
		s.Head.next = nil
		s.tail.next = nil
	} else {
		s.Head = &element[T]{ Data: val }
		s.Head.next = s.tail
		s.tail = s.Head
		s.tail.next = s.Head.next
	}
}

// Removes the top element from the stack 
func (s *stack[T])pop(){
	if s.Head == nil {
		fmt.Println("Pop failed!!\nNo more elements in the stack!! (stack underflow)")
		return
	}
	println("Popped:", s.Head.Data)
	s.Head = s.Head.next
	s.tail = s.Head
}

// Display the stack 
func (s *stack[T])Diaplay() {
	fmt.Printf("\nSTACK : ")
	if s.Head == nil{
		fmt.Printf("nil")
	}
	for i := s.Head ; i != nil ; i = i.next {
		fmt.Printf("%v  ", i.Data)
	}
	fmt.Printf("\n\n")
}

func main() {
	// Define stack 
	var Stack stack[int]
	// Push, Pop, Display 
	Stack.push(1)
	Stack.push(2)
	Stack.push(10)
	Stack.push(20)
	Stack.push(30)
	Stack.push(5)
	Stack.Diaplay()
	Stack.pop()
	Stack.pop()
	Stack.Diaplay()
	Stack.pop()
	Stack.pop()
	Stack.pop()
	Stack.pop()
	Stack.Diaplay()
	Stack.pop()
}