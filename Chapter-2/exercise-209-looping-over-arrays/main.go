package main

import (
	"fmt"
)

func main() {
	names := []string{"Jim", "Jane", "Joe", "Juns"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

}
