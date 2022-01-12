package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var signals = []string{"test"}

func main() {
	websites := []string{
		"https://google.com",
		"https://go.dev",
		"https://github.com",
	}
	for _, v := range websites {
		go getStatusCode(v)
		wg.Add(1)
	}
	wg.Wait()

	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()

	fmt.Printf("%d for endpoint %s\n", res.StatusCode, endpoint)
}
