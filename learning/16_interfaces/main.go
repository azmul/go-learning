package main

// Interfaces are named collections of method signatures.

import (
	"fmt"
	"math"
)

//  A basic interface for geometric shapes
type geometry interface {
	area() float64
	perimeter() float64
}

// we’ll implement this interface on rect and circle types
type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

// we implement geometry on rectangle
func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// // we implement geometry on circle
func (r circle) area() float64 {
	return math.Pi * r.radius * r.radius
}
func (r circle) perimeter() float64 {
	return 2 * math.Pi * r.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}
func main() {
	r := rectangle{width: 10, height: 20}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
