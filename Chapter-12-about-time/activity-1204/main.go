package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()

	fmt.Println("The current time: ", current.Format(time.ANSIC))

	seconds := time.Duration(21966 * time.Second)

	future := current.Add(seconds)

	fmt.Println("6 hours, 6 minutes and 6 seconds from now the time will be:", future)
}
