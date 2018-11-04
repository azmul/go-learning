// Variadic functions can be called with any number of trailing arguments.
// For example, fmt.Println is a common variadic function.
package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total) // [1 2] 3   [1 2 3] 6   [1 2 3 4 5 6] 21
}
func main() {

	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using func(slice...)
	nums := []int{1, 2, 3, 4, 5, 6}
	sum(nums...)
}
