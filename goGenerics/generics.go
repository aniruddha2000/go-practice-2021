package main

import (
	"fmt"
)

type GenericType interface {
	~float64 | int | string
}

type SuperFloat float64

func min[T GenericType](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	var sf SuperFloat = 0.4
	fmt.Println(min(sf, 1.2))
	fmt.Println(min("world ", "Hello"))
}
