package main

import (
	"fmt"
)

func add5Value(count int) {
	count += 5

	fmt.Println("add5Value  :", count)
}

func add5VPoint(count *int) {
	*count += 5

	fmt.Println("add5Point  :", *count)
}

func main() {
	var count int
	fmt.Println("Count  :", count)
	add5Value(count)

	fmt.Println("add5Value Post  :", count)

	add5VPoint(&count)

	fmt.Println("add5Point Post  :", count)

}
