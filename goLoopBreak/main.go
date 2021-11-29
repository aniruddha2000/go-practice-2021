package main

import "fmt"

func main() {
	fmt.Println("This is loop break")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}
	for _, day := range days {
		if day == "Tuesday" {
			continue
		}
		fmt.Printf("The days is %v\n", day)
	}
}
