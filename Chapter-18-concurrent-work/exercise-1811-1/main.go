package main

import (
	"context"
	"fmt"
	"time"
)

// controlling timeouts
// cancelling go routines
// and passing metadata across your Go application
func exampleTimeout() {
	ctx := context.Background()

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Called the API")
	case <-ctxWithTimeout.Done():
		fmt.Println("oh no timeout expired", ctxWithTimeout.Err())
		// do some logic to handle this
	}
}

func main() {
	exampleTimeout()
}
