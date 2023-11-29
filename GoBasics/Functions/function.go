package main

import(
	"fmt"
)

func main(){
	const num int = 4
	// function call 
	var res int = getSquare(num);
	fmt.Println("Square of", num, "is", res)

	add, sub, mul, div := calculate(4, 2)
	fmt.Printf("add : %v\nsub : %v\nmul : %v\ndiv : %v\n", add, sub, mul, div)

	// call sum 
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	// If you already have multiple args in a slice, 
	// apply them to a variadic function using func(slice...)
    sum(nums...)

	// call intSeq
	nextInt := intSeq()
	fmt.Println("\nNext Int")
	fmt.Println("next :", nextInt())
	fmt.Println("next :", nextInt())
	fmt.Println("next :", nextInt())

	// the state is unique to that particular function
	myNextInt := intSeq()
	fmt.Println("\nMy Next Int")
	fmt.Println("next :", myNextInt())
	fmt.Println("next :", myNextInt())

	// call fact 
	res = fact(5)
	println("\nFactorial of 5 is", res)

	// Closures can also be recursive, but this requires the closure 
	// to be declared with a typed var explicitly before it’s defined.
	var fib func(n int) int
    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n-1) + fib(n-2)
    }
    fmt.Println("\nfibonacci of 5 is", fib(5))
}

// this function returns the square of a number
func getSquare(number int) int {
	return number * number
}

// in 'go' a function can return multiple values
func calculate(a int, b int) (int, int, int, int) {
	// this function returns addition, substraction, multiplication, and division 
	return a + b, a - b, a * b, a / b
}

// prints the sum
// Variadic functions can be called with any number of trailing arguments.
func sum(nums...int) {
	fmt.Println("nums :",nums)
	var total int = 0

	// calculate total  
	for _, num := range nums {
		total += num
	}
	fmt.Println("sum is", total)
}


// 'Go' supports anonymous functions, which can form closures. 
// Anonymous functions are useful when you want to define a function inline without having to name it.
// This function intSeq returns another function, which we define anonymously in the body of intSeq. 
// The returned function closes over the variable i to form a closure.
func intSeq() func() int {
	var i int = 0
	return func() int {
		i++
		return i
	}
}

// recursion
// returns factorial of a number 
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}