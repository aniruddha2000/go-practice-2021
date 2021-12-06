package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main()  {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	inputString, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name is %v\n", inputString)
	fmt.Printf("Name is %T\n", inputString)
}
