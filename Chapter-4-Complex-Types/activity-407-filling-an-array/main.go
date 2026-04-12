package main

import "fmt"

func incrementArrKey(arr [10]int) [10]int {

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	return arr

}

func main() {
	var arr [10]int

	arr = incrementArrKey(arr)

	fmt.Println(arr)

}
