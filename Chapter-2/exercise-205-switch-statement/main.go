package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Monday

	switch dayBorn {
	case time.Monday:
		fmt.Println("Monday's child is fair of face")
	case time.Tuesday:
		fmt.Println("Tuesday's child is fair of grace")
	case time.Wednesday:
		fmt.Println("Wednesday's child is fair of woe")
	case time.Thursday:
		fmt.Println("Thursday's child is fair of go")
	case time.Friday:
		fmt.Println("Friday's child is fair of giving")
	case time.Saturday:
		fmt.Println("Saturday's child is fair of living")
	case time.Sunday:
		fmt.Println("Sunday's child is fair of blithe")
	default:
		fmt.Println("Error, day born not valid")
	}
}
