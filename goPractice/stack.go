package main

import (
	"fmt"
)

type Stack struct{
	Data []int
}

func (s *Stack) Empty() bool {
	return len(s.Data) == 0
}

func (s *Stack) Push(v int)  {
	s.Data = append(s.Data, v)
}

func (s *Stack) Pop() int {
	if !s.Empty() {
		v := s.Data[len(s.Data) - 1]
		s.Data = s.Data[:len(s.Data)-1]
		return v
	}
	return -1
}

func main() {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Pop())
}
