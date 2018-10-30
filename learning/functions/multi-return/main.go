// Go has built-in support for multiple return values
package main

import (
	"fmt"
)

// The (int, int) in this function signature shows that the function returns 2 ints
func vals() (int, int) {
	return 3, 7
}

func sumDivide(a, b int) (int, int) {
	sum := a + b
	divide := a - b
	return sum, divide
}
func main() {

	a, b := vals()
	fmt.Println(a, b) // 3, 7

	// want a subset of the returned values, use the blank identifier _
	_, c := vals()
	fmt.Println(c) // 7

	sum, divide := sumDivide(1, 2)
	fmt.Println(sum, divide) // 3 -1

}
