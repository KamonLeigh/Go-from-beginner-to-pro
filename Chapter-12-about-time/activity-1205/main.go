package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	newYork, err := time.LoadLocation("America/New_York")

	if err != nil {
		panic(err)
	}

	losAngeles, err := time.LoadLocation("America/Los_Angeles")

	if err != nil {
		panic(err)
	}

	jamaica, err := time.LoadLocation("America/Jamaica")

	if err != nil {
		panic(err)
	}

	fmt.Printf("The local current time is: %v \n", current.Format(time.ANSIC))
	fmt.Printf("The time in New York is: %v \n", current.In(newYork).Format(time.ANSIC))
	fmt.Printf("The time in Los Angeles %v \n", current.In(losAngeles).Format(time.ANSIC))
	fmt.Printf("The time in Jamaica %v \n", current.In(jamaica).Format(time.ANSIC))
}
