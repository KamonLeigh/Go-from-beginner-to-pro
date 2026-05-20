package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name provided")
	ErrInvalidRoutingNumber = errors.New("invalid routing number provided")
)

type directDeposit struct {
	lastName      string
	firtName      string
	bankName      string
	routingNumber int
	accountNumber int
}

func (d *directDeposit) validateRoutingNumber() error {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if d.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}

	return nil
}

func (d *directDeposit) validateLastName() error {
	if len(d.lastName) == 0 {
		return ErrInvalidLastName
	}

	return nil
}

func (d *directDeposit) customerDetails() {

	fmt.Printf("Last Name: %v\n", d.lastName)
	fmt.Printf("First Name: %v\n", d.firtName)
	fmt.Printf("Routing Number: %d\n", d.routingNumber)
	fmt.Printf("Account Number: %d\n", d.accountNumber)
}

func main() {
	customer := directDeposit{lastName: "", firtName: "Abe", bankName: "XYZ Inc", routingNumber: 17, accountNumber: 1809}

	err := customer.validateRoutingNumber()

	if err != nil {
		fmt.Println(err)
	}

	err = customer.validateLastName()
	if err != nil {
		fmt.Println(err)
	}

	customer.customerDetails()

}
