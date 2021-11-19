package main

import "fmt"

type Company struct {
	Name     string
	Location string
	Salary   string
}

func main() {
	UquinixMetal := Company{"UquinixMetal", "US", "180k"}
	fmt.Println(UquinixMetal)
	var OxideComputers = Company{"OxideComputers", "US", "175k"}
	fmt.Println(OxideComputers)
	var NewRelic = &Company{"NewRelic", "US", "115k"}
	fmt.Println(NewRelic)
}
