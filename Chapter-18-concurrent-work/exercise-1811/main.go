package main

import (
	"context"
	"log"
	"time"
)

func countNumbers(ctx context.Context, r chan int) {
	v := 0

	for {
		select {
		case <-ctx.Done():
			r <- v
			return
		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func main() {
	r := make(chan int)
	ctx := context.TODO()

	c1, stop := context.WithCancel(ctx)
	go countNumbers(c1, r)

	go func() {
		time.Sleep(time.Millisecond * 3)
		stop()
	}()

	v := <-r
	log.Println(v)
}
