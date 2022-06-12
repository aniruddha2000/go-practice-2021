package main

import (
	"fmt"
	"sync"
)

func main() {

	// Create an unbuffered channel.
	ch := make(chan int)

	// Create the WaitGroup and add a count
	// of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	// Launch the goroutine and handle Done.
	go func() {
		goroutine("first", ch)
		wg.Done()
	}()

	// Launch the goroutine and handle Done.
	go func() {
		goroutine("second", ch)
		wg.Done()
	}()

	// Send a value to start the counting.
	ch <- 1

	// Wait for the program to finish.
	wg.Wait()
}

// goroutine simulates sharing a value.
func goroutine(name string, ch chan int) {
	for {
		// Wait for the value to be sent.
		// If the channel was closed, return.
		val, ok := <-ch
		if !ok {
			fmt.Printf("goroutine closed %s\n", name)
			return
		}

		// Display the value.
		fmt.Printf("goroutine name : %s | inc %d\n", name, val)

		// Terminate when the value is 10.
		if val == 10 {
			close(ch)
			fmt.Printf("Value reached to 10 : goroutine closed %s\n", name)
			return
		}

		// Increment the value and send it
		// over the channel.
		ch <- (val + 1)
	}
}
