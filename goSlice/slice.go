package main

import (
	"fmt"
)

func main() {
	// var slice = []string{"Meta", "Apple", "Google"}
	// fmt.Println(slice)

	// slice = append(slice, "TCS", "CTS", "Amazon", "Netflix")
	// fmt.Println(slice)

	// fmt.Println(append(slice[1:]))

	// var myInts = []int{3, 9, 6, 4}
	// sort.Sort(sort.IntSlice(myInts))
	// fmt.Println(myInts)

	// var myString = make([]string, 3, 3)
	// fmt.Println(myString)

	// index := 3
	// slice = append(slice[:index], slice[index+2:]...)

	// fmt.Println(slice)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	// // Reverse a slice
	// reversed_slice := []string{}
	// for i := range slice {
	// 	n := slice[len(slice)-1-i]
	// 	reversed_slice = append(reversed_slice, n)
	// }
	// fmt.Println("Reverse a  slice :")
	// fmt.Println(reversed_slice)

	// Insert an element at specific position in slice
	fmt.Println("\nInsert an element at specific position in slice")

	array1 := []int{1, 5}
	array2 := []int{2, 3, 4}

	n := array1[1:]
	array1 = append(array1[:1], array2...)
	array1 = append(array1, n...)

	fmt.Println(array1)

	slice := []string{"Meta", "Apple", "Google", "Amazon", "Netflix"}
	// slice2 := []string{"Deepsource", "Groww"}
	// p := slice[2:]

	// fmt.Println("p->", p)
	// slice = append(slice[:2], slice2...)
	// fmt.Println("slice->", slice)
	// slice = append(slice, p...)
	// fmt.Println("slice->", slice)

	res := insert(slice, 2, "Deepsource")
	fmt.Println(res)
}

func insert(a []string, index int, value string) []string {
	if len(a) == index {
		return append(a, value)
	}

	fmt.Println("a[:index+1]->", a[:index+1])
	fmt.Println("a[index:]->", a[index:])

	a = append(a[:index+1], a[index:]...)
	fmt.Println(a)
	a[index] = value
	return a
}
