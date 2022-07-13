package main

import (
	"fmt"
)

type AStruct struct {
	a int
	// b map[string]int
}

type BStruct struct {
	a int
	// b map[string]int
}

func main() {
	fmt.Println("Hello World")
	var i int
	i += 1
	fmt.Println(i)

	aa := AStruct{
		a: 2,
		// b: map[string]int{"aa": 1, "bb": 2},
	}

	bb := AStruct{
		a: 2,
		// b: map[string]int{"aa": 1, "bb": 2},
	}

	if aa == bb {
		fmt.Println("heh")
	}
}
