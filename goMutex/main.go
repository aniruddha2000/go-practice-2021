package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func deposite(val int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance %d\n", val, balance)
	balance += val
	mutex.Unlock()
	wg.Done()
}

func withdrawl(val int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance %d\n", val, balance)
	balance -= val
	mutex.Unlock()
	wg.Done()
}

func main() {
	balance = 1000
	var wg sync.WaitGroup
	wg.Add(2)
	go withdrawl(200, &wg)
	go deposite(500, &wg)
	wg.Wait()
	fmt.Printf("Now the balance is, %d\n", balance)
}
