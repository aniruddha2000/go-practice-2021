package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Hella")
	for i := 0; i <= 10; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
		wg.Wait()
		fmt.Println(i)
	}
}
