package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// create a file & write into it
	file, _ := os.Create("file.txt")
	writer := io.Writer(file)
	n, err := writer.Write([]byte("Hello"))

	// Open the file
	file, _ = os.Open("file.txt")
	defer file.Close()

	reader := io.Reader(file)
	buffer := make([]byte, 1)
	n, err = reader.Read(buffer)
	fmt.Println(n, err, string(buffer))

	for {
		n, err = reader.Read(buffer)
		fmt.Println(n, err, string(buffer))
		if err != nil {
			break
		}
	}
}
