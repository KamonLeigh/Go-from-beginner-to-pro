package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// Same determines whether the trees

func Walk(t *tree.Tree, c chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, c)
	c <- t.Value
	fmt.Println("Value: ", t.Value)
	Walk(t.Right, c)

}

// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	result := true
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		Walk(t1, c1)
		close(c1)
	}()

	go func() {
		Walk(t2, c2)
		close(c2)
	}()

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if !ok1 && !ok2 {
			break
		}

		if ok1 != ok2 {
			result = false
			break
		}

		if v1 != v2 {
			result = false
			break
		}

	}

	return result
}

func main() {
	t1 := tree.New(10)
	t2 := tree.New(10)

	outcome := Same(t1, t2)

	fmt.Println(outcome)

}
