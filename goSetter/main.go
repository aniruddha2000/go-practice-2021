package main

import (
	"fmt"
)

type Student struct {
	Standard int
	Grades   []float64
	Age      int
}

func (s *Student) setAge(age int) {
	s.Age = age
}

func (s *Student) getYGPA() float64 {
	sum := 0.0
	for _, v := range s.Grades {
		sum += v
	}
	return sum / float64(len(s.Grades))
}

func main() {
	fmt.Println("This is setter example")
	aniruddha := Student{3, []float64{9.7, 9.8, 9.5}, 22}
	fmt.Println(aniruddha)
	aniruddha.setAge(23)
	fmt.Println(aniruddha)
	fmt.Println(aniruddha.getYGPA())
}
