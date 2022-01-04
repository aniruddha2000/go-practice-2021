package main

import (
	"fmt"
	"net/url"
)

const URL string = "https://aniruddha.com:8000/golang?video=urlhandle&editor=vscode"

func main() {
	fmt.Println("Hey it's the URL handlling part")

	res, err := url.Parse(URL)
	handleErr(err)
	// fmt.Println(res.Scheme)
	// fmt.Println(res.Host)
	// fmt.Println(res.Path)
	// fmt.Println(res.Port())
	// fmt.Println(res.RawQuery)

	query := res.Query()
	fmt.Printf("Type of the query %T\n", query)
	// fmt.Println(query["editor"])
	for k, v := range query {
		fmt.Printf("%v - %v\n", k, v)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "pkg.go.dev",
		Path:   "/doc",
	}
	fmt.Println(partsOfUrl.String())
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
