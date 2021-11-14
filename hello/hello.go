package main

import (
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func printPermissions(filename string) {
	info, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}

	mode := info.Mode()

	fmt.Print("Owner: ")
	for i := 1; i < 4; i++ {
		fmt.Print(string(mode.String()[i]))
	}

	fmt.Print("\nGroup: ")
	for i := 4; i < 7; i++ {
		fmt.Print(string(mode.String()[i]))
	}

	fmt.Print("\nOther: ")
	for i := 7; i < 10; i++ {
		fmt.Print(string(mode.String()[i]))
	}
}

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	names := []string{"Aniruddha", "Adhiraj", "Eshan"}
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	printPermissions("/home/aniruddha/Desktop/go-projects/gols/gols")
}
