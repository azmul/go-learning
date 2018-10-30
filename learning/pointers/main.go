package main

// Go has pointers, allowing you to pass references to values
// and records within your program
// A pointer holds the memory address of a value.

import (
	"fmt"
)

func zeroVal(ival int) {
	ival = 0
}

// zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
// The *iptr code in the function body then dereferences the pointer from its memory
// address to the current value at that address. Assigning a value to
// a dereferenced pointer changes the value at the referenced address.
func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	zeroVal(i)
	fmt.Println(i) // 1
	// The &i syntax gives the memory address of i
	fmt.Println(&i) // 0xc000016090

	zeroPtr(&i)
	fmt.Println(i)  // 0
	fmt.Println(&i) // 0xc000016090
}
