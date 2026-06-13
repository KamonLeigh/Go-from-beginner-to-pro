package main

import "fmt"

type person struct {
	lname  string
	age    int
	salary float64
}

func main() {
	fname := "Joe"
	grades := []int{100, 87, 67}
	stated := map[string]string{"KY": "Kentucky", "WV": "West Virgina", "VA": "Virgina"}

	p := person{lname: "Lincoln", age: 210, salary: 2500.00}

	fmt.Printf("fname value: %#v type: %T \n", fname, fname)
	fmt.Printf("grades value: %#v type: %T \n", grades, grades)
	fmt.Printf("stated value: %#v type: %T \n", stated, stated)
	fmt.Printf("p value: %#v type: %T \n", p, p)
}
