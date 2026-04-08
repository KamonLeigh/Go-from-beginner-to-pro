package main

import "fmt"

func main() {
	var salesTaxTotal float64

	salesTaxTotal += 0.99 * 0.075
	salesTaxTotal += 2.75 * 0.015
	salesTaxTotal += 0.87 * 0.02

	fmt.Println("Sales Tax Total: ", salesTaxTotal)

}
