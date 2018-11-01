package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and have type error, a built-in interface
func builtInError(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message
		return arg, errors.New("Will not work with this number")
	}
	return arg + 3, nil // A nil value in the error position indicates that there was no error
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d --- %s", e.arg, e.prob)
}

func customError(arg int) (int, error) {
	if arg == 42 {
		// we use &argError syntax to build a new struct,
		// supplying values for the two fields arg and prob
		return -1, &argError{arg, "Cannot work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, arg := range []int{7, 42} {
		if _, err := builtInError(arg); err != nil {
			fmt.Println("Successed", err)
		} else {
			fmt.Println("Failed", err)
		}
	}
	for _, arg := range []int{7, 42} {
		if _, err := customError(arg); err != nil {
			fmt.Println("successed", err)
		} else {
			fmt.Println("failed", err)
		}
	}
	// If you want to programmatically use the data in a custom error,
	// you’ll need to get the error as an instance of the custom error type via type assertion.
	_, e := customError(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
