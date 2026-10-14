package main

import "fmt"

func sum(data []int, c chan int) {
	sum := 0
	for _, i := range data {
		sum += i
	}

	c <- sum
}
func main() {
	ch := make(chan int)

	s := []int{4, 7, 9 - 8, 2, 6}
	go sum(s[len(s)/2:], ch)
	go sum(s[:len(s)/2], ch)

	x, y := <-ch, <-ch

	fmt.Println(x, y, x+y)
}
