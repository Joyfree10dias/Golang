package main

import (
	"errors"
	"fmt"
)

// Error checking 
func f1(arg int) (int, error) {
	if arg == 100 {
		return arg, errors.New("Can't use 100")
	}
	return arg, nil
}

// Custom errors
type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 100 {
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {
	// Call f1()
	arg, err := f1(101)
	if err == nil {
		fmt.Println("No error, can use", arg)
	} else {
		fmt.Println(err)
	}

	// Call f2() which implements custom errors 
	_, err = f2(100)
	fmt.Println(err)
	// If we want to access arg and prob 
	if val, e := err.(*argError) ; e {
		fmt.Println("Value :", val.arg, "and Error :",val.prob)
	}
}