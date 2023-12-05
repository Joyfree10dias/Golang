package main

import(
	"fmt"
)

// Declare Linked List
type LinkedList[T comparable] struct {
	Head *element[T]
}

// Declare element 
type element[T comparable] struct {
	Data T
	next *element[T]
}

// Add element at the end of the list 
func (list *LinkedList[T])append(val T) {
	if list.Head == nil {
		list.Head = &element[T]{ Data: val }
		list.Head.next = nil
	} else {
		var tail = list.Head
		for tail.next != nil {
			tail = tail.next
		}
		tail.next = &element[T]{ Data: val }
		tail = tail.next
		tail.next = nil
	}
}

// Add element in the beginning 
func (list *LinkedList[T])addAtBeg(val T){
	if list.Head == nil {
		list.Head = &element[T]{ Data: val }
		list.Head.next = nil
	} else {
		curr := list.Head
		list.Head = &element[T]{ Data: val }
		list.Head.next = curr
	}
}

// This function adds elements between Linked List 
func (list *LinkedList[T])addInBetween(val T, pos int) {
	temp := list.Head
	i := 0
	for i < pos - 2 && temp.next != nil {
		temp = temp.next
		i++
	}
	curr := temp
	curr = curr.next
	if pos < 2 || curr == nil {
		fmt.Println("Invalid position :", pos)
		return
	}
	node := &element[T]{ Data: val }
	temp.next = node
	node.next = curr
}

// Removes the desired element 
func (list *LinkedList[T])remove(val T) {
	curr := list.Head
	temp := list.Head
	var c int = 0
	for curr != nil {
		if curr.Data == val {
			fmt.Println(val, "is deleted!!")
			if curr == list.Head {
				list.Head = list.Head.next
				return
			} else if curr.next == nil {
				for i := 0 ; i < c - 1 ; i++ {
					temp = temp.next
				}
				curr = temp
				curr.next = nil
				return
			} else {
				for i := 0 ; i < c - 1 ; i++ {
					temp = temp.next
				}
				temp.next = curr.next
				curr.next = nil
				return
			}
		}
		curr = curr.next
		c++
	}
	fmt.Println(val, "is not present in the list")
}

// Display the list 
func (list *LinkedList[T])Display() {
	curr := list.Head
	fmt.Printf("\nLIST : ")
	for curr != nil {
		fmt.Printf("%v --> ",curr.Data)
		curr = curr.next
	}
	fmt.Printf("nil\n\n")
}

func main() {
	// Initialise Linked List 
	var linkedList = LinkedList[int]{}
	// Append, Delete, Display 
	linkedList.addAtBeg(5)
	linkedList.append(9)
	linkedList.append(10)
	linkedList.append(1)
	linkedList.append(4)
	linkedList.Display()
	linkedList.addAtBeg(3)
	linkedList.Display()
	linkedList.remove(3)
	linkedList.remove(10)
	linkedList.remove(1)
	linkedList.Display()
	linkedList.addInBetween(3, 2)
	linkedList.addInBetween(6, 3)
	linkedList.addInBetween(11, 5)
	linkedList.addInBetween(23, 10)
	linkedList.addInBetween(12, 1)
	linkedList.addInBetween(13, 2)
	linkedList.addInBetween(14, 7)
	linkedList.addInBetween(145, 9)
	linkedList.addInBetween(123, 8)
	linkedList.Display()
}