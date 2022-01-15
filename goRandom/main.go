package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	// "math/rand"
	// "time"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	myRandom, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandom)
}
