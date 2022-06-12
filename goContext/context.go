package main

import (
	"context"
	"fmt"
	"time"
)

func contextSetter(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "key", "value")
	return ctx
}

func doSomething(ctx context.Context) {
	val := ctx.Value("key")
	fmt.Println(val)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timed Out")
			return
		default:
			fmt.Println("Doing Something")
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx = contextSetter(ctx)
	go doSomething(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("I have exited")
	}
	time.Sleep(2 * time.Second)
}
