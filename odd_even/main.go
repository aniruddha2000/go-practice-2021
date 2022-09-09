package main

import (
	"fmt"
)

func main() {

	oddCh := make(chan int)
	evenCh := make(chan int)

	go print(oddCh, evenCh)

	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			oddCh <- i
		} else {
			evenCh <- i
		}
	}
}

func print(odd <-chan int, eve <-chan int) {
	for {
		select {
		case v := <-odd:
			fmt.Println(v)
		case v := <-eve:
			fmt.Println(v)
		}
	}
}
