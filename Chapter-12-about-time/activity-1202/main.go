package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t := time.Date(2023, time.January, 31, 02, 49, 21, 0, time.Now().UTC().Location())

	fmt.Println(t)

	day := strconv.Itoa(t.Day())
	month := strconv.Itoa(int(t.Month()))
	year := strconv.Itoa(t.Year())

	minute := strconv.Itoa(t.Minute())
	hour := strconv.Itoa(t.Hour())
	second := strconv.Itoa(t.Second())

	fmt.Println(hour + ":" + minute + ":" + second + " " + year + "/" + month + "/" + day)

}
