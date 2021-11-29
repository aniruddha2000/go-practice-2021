package main

import "fmt"

func main() {
	var aFloat float64 = 87.9
	pointer := &aFloat

	*pointer = *pointer * 2
	fmt.Println(aFloat, pointer, *pointer)
}
