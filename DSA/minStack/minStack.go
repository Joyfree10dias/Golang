package main

import (
	"errors"
	"fmt"
)

// Min-Stack
type minStack struct {
	stack []int32
	topRef int32 // Points to the top of the stack
	minValStack []int32 // Stores the min values of the stack
} 

// Intialise Min-Stack 
func (ms *minStack)init() {
	ms.topRef = -1
}

// Push element inside Min-Stack
func (ms *minStack)push(data int32) {
	// Check for min value 
	if ms.minValStack == nil {
		ms.minValStack = append(ms.minValStack, data)
	} else if data < ms.stack[ms.topRef] {
		ms.minValStack = append(ms.minValStack, data)
	} else if data > ms.stack[ms.topRef] {
		ms.minValStack = append(ms.minValStack, ms.minValStack[ms.topRef])
	}
	ms.topRef++
	// ms.stack[ms.topRef] = data
	ms.stack = append(ms.stack, data)
	fmt.Println("Pushed :", data)
}

// Pop element
func (ms *minStack)pop() {
	if ms.topRef != -1 {
		fmt.Println("Popped :", ms.stack[ms.topRef])
		ms.stack = ms.stack[: ms.topRef]
		ms.minValStack = ms.minValStack[: ms.topRef]
		ms.topRef--
	} else {
		fmt.Println("Stack empty!! Nothing to pop.")
	}
}

// This func gets the top element of the stack.
func (ms *minStack)top() (int32, error) {
	if ms.topRef == -1 {
		return -1, errors.New("stack empty")
	} else {
	return ms.stack[ms.topRef], nil
	}
}

// This func retrieves the minimum element in the stack.
func (ms *minStack)getMin() int32 {
	return ms.minValStack[ms.topRef]
}

// To display Stack 
func (ms *minStack)display() {
	var i int32
	fmt.Printf("\nDisplay : ")
	if ms.topRef == -1 {
		fmt.Println("Stack empty!!")
	} else {
		for i = ms.topRef ; i > 0 ; i-- {
			fmt.Printf("%d --> ", ms.stack[i])
		}
		fmt.Printf("%d \n\n", ms.stack[i])
	}
}


func main() {
	// Declare a Min Stack 
	var ms minStack
	ms.init()
	ms.push(23)
	ms.push(2)
	ms.push(11)
	ms.push(13)
	ms.display()
	ms.pop()
	ms.push(50)
	ms.pop()
	ms.push(-6)
	ms.display()
	topVal, topErr := ms.top()
	if topErr != nil {
		fmt.Println("Error :", topErr)
	} else {
		fmt.Println("Top :", topVal)
	}
	min := ms.getMin()
	fmt.Println("Min value of the stack :", min)
	ms.pop()
	ms.display()
	min = ms.getMin()
	fmt.Println("Min value of the stack :", min)
}