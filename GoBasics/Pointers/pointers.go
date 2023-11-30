package main

import(
	"fmt"
)

func zeroVal(val int){
	fmt.Println("address of val :",&val)
	val = 0
}

func zeroPtr(ptr *int){
	fmt.Println("ptr :",ptr)
	*ptr = 0
}

func main(){
	i := 1
	fmt.Println("initial :",i)

	zeroVal(i)
	fmt.Println("zeroVal :",i)

	zeroPtr(&i)
	fmt.Println("zeroPtr :",i)

	fmt.Println("address of i :",&i)

}