package main

import "fmt"

func main() {
	company := make(map[string]string)
	company["Google"] = "120k"
	company["Netflix"] = "125k"
	company["Amazon"] = "110k"
	company["TCS"] = "40k"
	fmt.Println(company)
	fmt.Println(company["Amazon"])

	for key, value := range company {
		fmt.Printf("Salary of %v is %v\n", key, value)
	}

	delete(company, "TCS")
	for key, value := range company {
		fmt.Printf("Salary of %v is %v\n", key, value)
	}
}
