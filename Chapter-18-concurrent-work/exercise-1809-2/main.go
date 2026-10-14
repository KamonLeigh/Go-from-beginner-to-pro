package main

import "log"

func doSomething() chan int {
	ch := make(chan int)

	go func() {
		for i := range ch {
			log.Println(i)
		}
	}()

	return ch
}
func main() {
	ch := doSomething()

	ch <- 1
	ch <- 4
	close(ch)
}
