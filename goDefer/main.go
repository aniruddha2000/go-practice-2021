package main

import "fmt"

func main() {
	defer fmt.Println("Hehe")
	defer fmt.Println("Defer example")
	fmt.Println("This is")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
