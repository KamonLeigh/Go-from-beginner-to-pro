package main

import "fmt"

func checkNumbers(i int) (int, string) {
	switch {
	case i%2 == 0:
		return i, "Even"
	default:
		return i, "Odd"
	}
}

func main() {
	for i := 0; i <= 15; i++ {
		num, result := checkNumbers(i)

		fmt.Printf("Results %d %s\n", num, result)
	}
}
