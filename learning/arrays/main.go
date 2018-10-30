package main

import (
	"fmt"
)

func main() {

	// Initialize an array
	var a [5]int   // By default an array is zero-valued
	fmt.Println(a) // [0 0 0 0 0]

	//
	a[4] = 100        // We can set a value at an index using the array[index] = value syntax
	fmt.Println(a)    // [0 0 0 0 100]
	fmt.Println(a[4]) // 100

	fmt.Println(len(a)) // The builtin len returns the length of an array.

	b := [5]int{1, 2, 3, 4, 5} // initialize an array in one line
	fmt.Println(b)             // [1 2 3 4 5]

	var two [2][3]int // multi-dimensional data structures
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			two[i][j] = i + j
		}
	}
	fmt.Println(two)
}
