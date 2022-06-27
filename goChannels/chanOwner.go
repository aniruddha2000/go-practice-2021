package main

import (
	"fmt"
	// "time"
)

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			defer fmt.Println("Done sending!")
			for i := 0; i <= 5; i++ {
				fmt.Printf("Sending: %d\n", i)
				resultStream <- i
				// time.Sleep(1 * time.Millisecond)
			}
		}()
		return resultStream
	}
	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
