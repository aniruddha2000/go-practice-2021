package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://aniruddha2000.github.io/aniruddhabasak/"

func main() {
	fmt.Println("It's the we request part")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	databyte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))
}
