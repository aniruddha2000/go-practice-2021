package main

import "fmt"

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	repeatStream := make(chan interface{})
	go func() {
		defer close(repeatStream)
		for {
			for v := range values {
				select {
				case <-done:
					return
				case repeatStream <- v:
				}
			}
		}
	}()
	return repeatStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	for num := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
}
