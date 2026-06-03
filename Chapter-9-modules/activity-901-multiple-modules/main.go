package main

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	id := uuid.New()
	quote := quote.Opt()

	fmt.Printf("Generated UUID: %s \n", id)
	fmt.Printf("Random Quote: %s \n", quote)

}
