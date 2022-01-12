package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	mut.RLock()
	var score = []int{0}
	mut.RUnlock()

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("First routine")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Second routine")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Third routine")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Fourth routine")
		m.RLock()
		fmt.Println(score)
		m.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
