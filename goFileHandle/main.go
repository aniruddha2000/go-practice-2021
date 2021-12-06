package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("This is file handlling in go")
	content := "This is Money Heist S05"

	file, err := os.Create("./myMoneyHeist.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println(length)
	readFile("./myMoneyHeist.txt")
}

func readFile(name string) {
	dataByte, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataByte))

}
