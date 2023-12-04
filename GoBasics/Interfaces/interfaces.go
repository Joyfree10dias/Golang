package main

import(
	"fmt"
)

// Declare Interfaces 
type geometry interface {
	area() float32
	perimeter() float32
}

// Declare rectangle
type rectangle struct {
	length, breadth float32
}

// Declare circle 
type circle struct {
	radius float32
}

// Declare all the methods related to both structs
// Calculate area for rectangle  
func (r rectangle) area() float32 {
	return r.length * r.breadth
}

// Calculate perimeter for rectangle  
func (r rectangle) perimeter() float32 {
	return 2 * r.length + 2 * r.breadth
}

// Calculate area for circle  
func (c circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

// Calculate perimeter for circle  
func (c circle) perimeter() float32 {
	return 2 * 3.14 * c.radius
}

// call area and perimeter using geometry interface
// This function prints area and perimeter 
func measure(g geometry) {
	fmt.Println("Struct :",g)
	fmt.Println("Area :", g.area())
	fmt.Println("Perimeter :",g.perimeter())
}

// Struct Embedding
type base struct {
	num int
}

func (b base) printNumber() {
	fmt.Println("Number :",b.num)
}

type container struct {
	base 
	str string
}

func main() {
	// define both structs 
	// rectangle
	r := rectangle{ length: 10, breadth: 10 }
	
	// circle 
	c := circle{ radius: 4 }

	// Pass both r and c to measure
	measure(r)
	fmt.Println()
	measure(c)

	// Define a container 
	co := container {
		base : base{
			num : 12,
		},
		str : "Here is a string",
	}

	// We can access the base’s fields directly on co, e.g. co.num.
	fmt.Printf("\nco={ num: %v, str: %v }\n", co.num, co.str)
	// Alternatively, we can spell out the full path using the embedded type name.
	fmt.Println("also num :", co.base.num)
	co.printNumber()

	// Declare interface
	type printer interface {
        printNumber()
    }

	fmt.Println("\nPrinter:")
	var p printer = co
	p.printNumber()


}