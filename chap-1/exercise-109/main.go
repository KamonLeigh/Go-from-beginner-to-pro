package main

import "fmt"

func main() {
	// Main course
	var total float64 = 2 * 13
	println("Sub :", total)

	// Drinks
	total = total + (4 * 2.25)
	println("Sub :", total)

	// Discount
	total = total - 5
	println("Sub :", total)

	// 10% tip
	tip := total * 0.1
	println("Tip:", tip)

	total = total + tip
	println("Total :", total)
	// split bill
	split := total / 2
	println("Split :", split)

	// Reward every 5th visit
	visitCount := 24
	visitCount = visitCount + 1

	remainder := visitCount % 5

	if remainder == 0 {
		fmt.Println("With this visit, you've earned a reward.")
	}

}
