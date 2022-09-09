package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("vim-go")

	ch := make(chan int, 2)

	wg.Add(1)
	go func(ch chan int) {
		for i := 0; i < 3; i++ {
			ch <- i
			ch <- i
		}
		wg.Done()
	}(ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
