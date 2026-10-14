package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
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

func HelloServer(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("API Response!")
	case <-ctx.Done():
		fmt.Println("oh no the context expired")
		return
	}
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
