package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan string, 9)

	go func() {
		for p := range ch {
			fmt.Println("employee : working :", p)
			wg.Done()
		}
	}()

	const work = 10
	for w := 0; w < work; w++ {
		wg.Add(1)
		ch <- fmt.Sprintf("Paper %v", w)
	}

	wg.Wait()
	close(ch)
}
