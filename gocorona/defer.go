package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	b()
}
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}
