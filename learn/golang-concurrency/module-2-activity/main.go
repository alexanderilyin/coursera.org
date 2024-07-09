package main

// Usage: go run -race .

// Two goroutines write to counter variable w/o propper synchronization
// which makes possible for some of the increments be lost

import (
	"fmt"
	"time"
)

func main() {
	var counter int

	// First goroutine
	go func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
	}()

	// Second goroutine
	go func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
	}()

	// Wait for goroutines to finish
	time.Sleep(time.Second)

	fmt.Println("Final counter value:", counter)
}
