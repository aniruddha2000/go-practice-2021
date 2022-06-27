package main

import (
	"bytes"
	"fmt"
	"os"
	// "time"
)

func main() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)
	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		// defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		defer fmt.Println("Producer Done.")
		for i := 0; i < 5; i++ {
			// fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			fmt.Printf("Sending: %d\n", i)
			intStream <- i
			// time.Sleep(1 * time.Millisecond)
		}
	}()
	for integer := range intStream {
		// fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
		fmt.Printf("Received %v.\n", integer)
	}
}
