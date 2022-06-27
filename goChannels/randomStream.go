package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func newRandStream(done <-chan interface{}, wg *sync.WaitGroup) <-chan int {
	randStream := make(chan int)

	go func() {
		defer wg.Done()
		defer fmt.Println("randStream exited")
		defer close(randStream)
		for {
			select {
			case randStream <- rand.Int():
			case <-done:
				return
			}
		}
	}()
	return randStream
}

func main() {
	var wg sync.WaitGroup

	done := make(chan interface{})

	wg.Add(1)
	randStream := newRandStream(done, &wg)
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	wg.Wait()
}
