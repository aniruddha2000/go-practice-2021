package main

import (
	"encoding/json"
	"fmt"
)

type Company struct {
	Name     string `json:"coursename"`
	Location string `json:"location"`
	Salary   string `json:"salary"`
}

func main() {
	fmt.Println("This is JSON handle")
	// EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {
	companies := []Company{
		{"Google", "UK", "150k"},
		{"Netflix", "US", "174k"},
		{"Microsoft", "US", "120k"},
	}

	finalJSON, err := json.MarshalIndent(companies, "", "\t")
	handleErr(err)
	fmt.Printf("%s\n", finalJSON)
}

func DecodeJSON() {
	jsonData := []byte(`
	{
		"coursename": "Google",
		"location": "UK",
		"salary": "150k"
	}
	`)

	var company Company

	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("Valid json!")
		json.Unmarshal(jsonData, &company)
		fmt.Printf("%#v\n", company)
	} else {
		fmt.Println("JSON NOT VALID!")
	}

	// some case where you just want to add data to key value

	var myCompanyData map[string]interface{}
	json.Unmarshal(jsonData, &myCompanyData)
	fmt.Printf("%#v\n", myCompanyData)

	for k, v := range myCompanyData {
		fmt.Printf("Key = %v, value is %v, Type is %T\n", k, v, v)
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
