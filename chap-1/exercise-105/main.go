package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	Debug, LogLevel, StartUpTime := getConfig()

	fmt.Println(Debug, LogLevel, StartUpTime)
}
