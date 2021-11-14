package main

import "fmt"

func main() {
	var numbers [3]string
	numbers[0] = "zero"
	numbers[1] = "One"
	numbers[2] = "Two"

	fmt.Println(numbers)

	var colors = [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors)
}
