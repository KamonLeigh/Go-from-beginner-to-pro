package main

import (
	"fmt"
)

var elements = []string{
	"Good",
	"Good",
	"Bad",
	"Good",
	"Good",
}

func deleteBadElement() []string {
	result := elements[:0:0]
	for _, value := range elements {
		if value != "Bad" {
			result = append(result, value)
		}
	}

	return result
}

func main() {

	elements := deleteBadElement()
	fmt.Println("Elements", elements)
}
