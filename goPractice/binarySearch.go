package main

import "fmt"

func binarySearch(find int, items []int) bool {
	low := 0
	high := len(items) - 1

	for low <= high {
		median := (low + high) / 2

		if items[median] < find {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(items) || items[low] != find {
		return false
	}
	return true
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items))
}
