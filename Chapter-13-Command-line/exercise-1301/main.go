package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <name>")
	}

	name := args[1]

	greeting := fmt.Sprintf("Hello, %s! Welcome to the command line", name)

	fmt.Printf(greeting)
}
