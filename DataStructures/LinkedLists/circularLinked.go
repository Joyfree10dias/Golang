package main

import(
	"fmt"
)

// Declare Linked List
type LinkedList[T comparable] struct {
	Head, tail *element[T]
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
		list.Head.next = list.Head
		list.tail = list.Head
		list.tail.next = list.tail
	} else {
		list.tail.next = &element[T]{ Data: val }
		list.tail = list.tail.next
		list.tail.next = list.Head
	}
}

// Add element in the beginning 
func (list *LinkedList[T])addAtBeg(val T){
	if list.Head == nil {
		list.Head = &element[T]{ Data: val }
		list.Head.next = list.Head
		list.tail = list.Head
		list.tail.next = list.tail
	} else {
		curr := list.Head
		list.Head = &element[T]{ Data: val }
		list.Head.next = curr
		list.tail.next = list.Head
	}
}

// This function adds elements between Linked List 
func (list *LinkedList[T])addInBetween(val T, pos int) {
	temp := list.Head
	i := 0
	for i < pos - 2 && temp.next != list.Head {
		temp = temp.next
		i++
	}
	curr := temp
	curr = curr.next
	if pos < 2 || curr == list.Head {
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
				list.tail.next = list.Head
				return
			} else if curr.next == list.Head {
				for i := 0 ; i < c - 1 ; i++ {
					temp = temp.next
				}
				curr = temp
				curr.next = list.Head
				list.tail = curr
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
	for curr.next != list.Head {
		fmt.Printf("%v --> ",curr.Data)
		curr = curr.next
	}
	fmt.Printf("%v --> ", curr.Data)
	curr = curr.next
	fmt.Printf("%v\n\n",curr.Data)
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
	linkedList.remove(4)
	linkedList.remove(10)
	linkedList.addAtBeg(69)
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
	linkedList.append(8)
	linkedList.Display()
	
}