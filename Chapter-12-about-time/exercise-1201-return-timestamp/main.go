package main

import (
	"fmt"
	"time"
)

func whatistheclock() string {
	return time.Now().Format(time.ANSIC)
}

func main() {
	fmt.Println(whatistheclock())
}
