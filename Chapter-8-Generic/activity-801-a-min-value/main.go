package main

import "fmt"

func findMinGeneric[num float64 | int](values []num) num {
	if len(values) == 0 {
		return -1
	}

	min := values[0]

	for _, value := range values {
		if min > value {
			min = value
		}
	}

	return min
}

func main() {
	slice := []int{1, 32, 5, 8, 10, 11}

	fmt.Printf("Min value is: %v\n", findMinGeneric(slice))

	slice2 := []float64{1.1, 32.1, 5.1, 8.1, 10.1, 11.1}

	fmt.Printf("Min value is %v\n", findMinGeneric(slice2))

}
