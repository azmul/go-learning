package main

// We can use channels to synchronize execution across goroutines.
// Here’s an example of using a blocking receive to wait for
// a goroutine to finish.

import (
	"fmt"
	"time"
)

// This is the function we’ll run in a goroutine.
// The done channel will be used to notify another goroutine that
// this function’s work is done.
func worker(done chan bool) {

	fmt.Printf("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true // Send a value to notify that we’re done.

}

func main() {
	done := make(chan bool, 2)
	go worker(done) // Start a worker goroutine, giving it the channel to notify on.

	<-done // Block until we receive a notification from the worker on the channel.
	// Note: If you removed the <- done line from this program,
	// the program would exit before the worker even started.
}
