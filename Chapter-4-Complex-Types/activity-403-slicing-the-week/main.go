package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	temp := []string{days[len(days)-1]}
	days = append(temp, days[0:(len(days)-1)]...)
	fmt.Println(days)

}
