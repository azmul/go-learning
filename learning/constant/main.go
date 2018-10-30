package main

import (
	"fmt"
	"math"
)

// const declares a constant value.

const s string = "constant"

func main() {
	fmt.Println(s)

	// const n int = 500000  // cannot use n (type int) as type float64 in argument to math.Sin
	const n = 500000 // A numeric constant has no type until it’s given one, such as by an explicit cast
	const d = 3e10 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	// fmt.Println(math.Sin(n))
	fmt.Println(math.Sin(n))

}
