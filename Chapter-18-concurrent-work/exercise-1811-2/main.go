package main

import (
	"context"
	"fmt"
)

// controlling timeouts
// cancelling go routines
// and passing metadata across your Go application

func exampleWithValues() {
	ctx := context.Background()
	ctxWithValue := context.WithValue(ctx, "user", "123")

	if userID, ok := ctxWithValue.Value("user").(string); ok {
		fmt.Println("this is the userID", userID)
	} else {
		fmt.Println("this is a protected route - no userID found")
	}
}

func main() {
	exampleWithValues()
}
