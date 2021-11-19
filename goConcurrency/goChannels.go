package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sendValue(c chan int, val int) {
	defer wg.Done()
	c <- val * 5
}

func main() {
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sendValue(c, i)
	}
	wg.Wait()
	close(c)
	for item := range c {
		fmt.Println(item)
	}
}
