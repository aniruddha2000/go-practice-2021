package main

import "fmt"

type Queue struct {
	Data []int
}

func (q *Queue) Empty() bool {
	return len(q.Data) == 0
}

func (q *Queue) Enqueue(v int) {
	q.Data = append(q.Data, v)
}

func (q *Queue) Dequeue() int {
	if !q.Empty() {
		v := q.Data[0]
		q.Data = q.Data[1:len(q.Data)]
		return v
	}
	return -1
}

func main() {
	q := &Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	fmt.Println(q.Dequeue())
}
