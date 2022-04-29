package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

type Config struct {
	a []int
}

func main() {
	var (
		wg sync.WaitGroup
		v  atomic.Value
	)

	v.Store(Config{a: []int{}})

	go func() {
		var i int
		for {
			i++
			cfg := Config{
				a: []int{i + 1, i + 2, i + 3, i + 4, i + 5},
			}
			v.Store(cfg)
		}
	}()

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			cfg, ok := v.Load().(Config)
			if !ok {
				log.Fatalf("Received a different config %T", cfg)
			}
			fmt.Println("Config ", cfg)
		}()
	}

	wg.Wait()
}
