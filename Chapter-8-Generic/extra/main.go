package main

import "fmt"

type Number interface {
	int | float32 | float64
}

type Price[T Number, K string] struct {
	Money T
	Title K
}

func (p Price[T, K]) Add(increase T) Price[T, K] {
	p.Money += increase
	return p
}

func main() {
	foo := Price[int, string]{
		Money: 78,
		Title: "fix",
	}

	bar := Price[int, string]{
		Money: 89,
		Title: "No Diggity",
	}

	fmt.Println(foo)
	fmt.Println(bar)

}
