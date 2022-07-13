package main

import "fmt"

func main() {
	index := 2
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = 10
	fmt.Println(slice)
}
