package main

import (
	"fmt"
	"sort"
)

func main() {
	var slice = []string{"Meta", "Apple", "Google"}
	fmt.Println(slice)

	slice = append(slice, "Amazon", "Netflix")
	fmt.Println(slice)

	fmt.Println(append(slice[1:]))

	var myInts = []int{3, 9, 6, 4}
	sort.Sort(sort.IntSlice(myInts))
	fmt.Println(myInts)

	var myString = make([]string, 3, 3)
	fmt.Println(myString)
}
