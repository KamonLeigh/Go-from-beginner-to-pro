package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name provided")
	ErrInvalidRoutingNumber = errors.New("invalid routing number provided")
)

func main() {
	fmt.Println(ErrInvalidLastName)
	fmt.Println(ErrInvalidRoutingNumber)
}
