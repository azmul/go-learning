package main

import (
	"fmt"
)

// Go supports methods defined on struct types.

type reactangle struct {
	width, height int
}

// This area method has a receiver type of *reactangle
func (r *reactangle) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types
func (r reactangle) perimeter() int {
	return 2*r.width + 2*r.height
}
func main() {
	r := reactangle{width: 10, height: 20}

	// we call the 2 methods defined for our struct
	fmt.Println(r.area())      // 200
	fmt.Println(r.perimeter()) // 60

	// Go automatically handles conversion between values and pointers for method calls
	rp := &r
	fmt.Println(rp.area())      // 200
	fmt.Println(rp.perimeter()) // 60
}
