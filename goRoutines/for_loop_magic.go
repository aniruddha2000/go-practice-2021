package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	for _, v := range []string{"Hello", "world", "Aniruddha"} {
		wg.Add(1)
		go func (v string)  {
			fmt.Println(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
