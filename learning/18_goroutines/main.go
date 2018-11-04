package main

import (
	"fmt"
)

// var w sync.WaitGroup

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, "", i)
	}
	// w.Done()
}

func main() {
	// w.Add(2)

	// we’d call that in the usual way, running it synchronously
	f("direct")

	// To invoke this function in a goroutine, use go f(s).
	// This new goroutine will execute concurrently with the calling one
	go f("goroutine")

	// w.Wait()

	// Can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// This Scanln requires we press a key before the program exits
	fmt.Scanln()
	fmt.Println("done")
}
