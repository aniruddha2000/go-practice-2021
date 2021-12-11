package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func main() {
	fmt.Println("vim-go")
	v := MyFloat(-math.Sqrt2)
	fmt.Println(v)
}
