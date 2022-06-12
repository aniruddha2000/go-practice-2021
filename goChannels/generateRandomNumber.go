package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Declare constant for number of goroutines.
const max = 100

func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the buffered channel with room for
	// each goroutine to be created.
	ch := make(chan int, max)

	var wg sync.WaitGroup

	// Iterate and launch each goroutine.
	wg.Add(max)
	for i := 0; i < max; i++ {
		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
		go func() {
			ch <- rand.Intn(1000)
			wg.Done()
		}()
	}

	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.

	go func() {
		wg.Wait()
		close(ch)
	}()

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints.
	store := []int{}
	for val := range ch {
		store = append(store, val)
	}

	// Print the values in our slice.
	fmt.Println(store)
}
