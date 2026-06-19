package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	current := time.Now()

	day := strconv.Itoa(current.Day())
	month := strconv.Itoa(int(current.Month()))
	year := strconv.Itoa(current.Year())
	hour := strconv.Itoa(current.Hour())
	minute := strconv.Itoa(current.Minute())
	seconds := strconv.Itoa(current.Second())

	fmt.Println(hour + ":" + minute + ":" + seconds + " " + year + "/" + month + "/" + day)
}
