package main

import "fmt"

func main() {
	var aFloat float64 = 87.9
	pointer := &aFloat

	fmt.Println(aFloat, pointer)
}
