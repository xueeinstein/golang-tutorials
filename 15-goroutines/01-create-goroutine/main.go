package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for i := 0; i < 5; i++ {
		// Uncomment to disable sleep in goroutine, we wouldn't see "Welcome", because it executes instantly
		time.Sleep(100 * time.Microsecond)

		fmt.Println(str)
	}
}

func main() {
	fmt.Println(time.Microsecond)
	// Calling Goroutine with `go` keyword
	go display("Welcome")

	// Calling normal function
	display("Hello")

	// Advangate of Goroutine
	// 1. cheaper than threads
	// 2. are stored in stack and its size can grow and shrink correspondingly, whereas stack of thread fixed.
	// 3. Goroutines can communicate using the channel and these channels are specially designed to prevent race conditions when accessing shared memory using Goroutines
	// 4. Suppose a program has one thread, and that thread has many Goroutines associated with it. If any of Goroutine blocks the thread due to resource requirement then all the remaining Goroutines will assign to a newly created OS thread. All these details are hidden from the programmers.

}
