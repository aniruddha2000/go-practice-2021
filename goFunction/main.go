package main

import "fmt"

func main() {
	var a, b int
	a, b = 5, 5
	fmt.Printf("From the main: a: %v, b: %v\n", a, b)
	adder(a, b)
	fmt.Printf("From the main: a: %v, b: %v\n", a, b)

	add, msg := proAdder(2, 3, 4, 6)
	fmt.Printf("add: %v\n", add)
	fmt.Printf("msg: %v\n", msg)
}

func adder(a int, b int) {
	c := a + b
	fmt.Printf("From the adder: c: %v\n", c)
}

func proAdder(values ...int) (int, string) {
	res := 0

	for _, val := range values {
		res += val
	}

	msg := "Here we added all your values"
	return res, msg
}
