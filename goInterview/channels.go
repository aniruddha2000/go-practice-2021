package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(i int, ch chan int) {
	defer wg.Done()
	ch <- i
}

// func bar(ch chan int) {
// 	for i := 0; i < 10; i++ {
// 		foo(i, ch)
// 	}
// 	close(ch)
// }

func main() {
	ch := make(chan int)

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go foo(i, ch)
	}

	// go bar(ch)
	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}

}
