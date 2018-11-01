// Functions are central in Go
package main

import (
	"fmt"
)

// it takes two ints and returns their sum as an int
func plus(a, b int) int {
	return a + b
}

// it takes three ints and returns their sum as an int
func plusPlus(a, b, c int) int {
	return a + b + c
}
func main() {
	twoSum := plus(1, 2)
	fmt.Println(twoSum) // 3

	threeSum := plusPlus(1, 2, 3)
	fmt.Println(threeSum) // 6
}
