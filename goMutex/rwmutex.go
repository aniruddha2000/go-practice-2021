package main

import (
	"fmt"
	"sync"
	"time"
)

type SyncData struct {
	lock sync.RWMutex
	wg   sync.WaitGroup
}

func main() {
	// m := map[int]int{}

	var sc SyncData

	sc.wg.Add(7)

	go readLoop(&sc)
	go readLoop(&sc)
	go readLoop(&sc)
	go readLoop(&sc)
	go writeLoop(&sc)
	go writeLoop(&sc)
	go writeLoop(&sc)

	sc.wg.Wait()
}

func writeLoop(sc *SyncData) {
	sc.lock.Lock()
	time.Sleep(1 * time.Second)
	fmt.Println("Write lock")
	fmt.Println("Write unlock")
	sc.lock.Unlock()
	sc.wg.Done()
}

func readLoop(sc *SyncData) {
	sc.lock.RLock()
	time.Sleep(1 * time.Second)
	fmt.Println("Read lock")
	fmt.Println("Read unlock")
	sc.lock.RUnlock()
	sc.wg.Done()
}
