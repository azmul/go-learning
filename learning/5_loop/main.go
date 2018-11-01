package main

import "fmt"

func main() {
	// for is Go’s only looping construct. Here are three basic types of for loops.
	i := 1
	for i <= 3 { // The most basic type, with a single condition like while loop
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j <= 5; j++ { // Common loop
		fmt.Println(j)
	}

	for { // infinte loop
		fmt.Println("loop")
		break
	}
	for { // infinte loop
		fmt.Println("loop")
		return
	}
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
