package main

import "fmt"

type Person struct {
	First string
	Last  string
}

type SecretAgent struct {
	Person Person
	Org    string
}

type Human interface {
	speak()
}

func (s SecretAgent) speak() {
	fmt.Println("I am secret agent ", s.Person.First)
}

func (p Person) speak() {
	fmt.Println("I am Person ", p.First)
}

func American(h Human) {
	fmt.Println("I am american as well as human")
	h.speak()
}

func main() {
	sa1 := SecretAgent{
		Person: Person{First: "James", Last: "Bond"},
		Org:    "CIA",
	}
	sa2 := SecretAgent{
		Person: Person{First: "Ajit", Last: "Doval"},
		Org:    "RAW",
	}

	p1 := Person{First: "Dr.", Last: "Strange"}
	// sa1.speak()
	// sa2.speak()
	// p1.speak()

	American(sa1)
	American(sa2)
	American(p1)
}
