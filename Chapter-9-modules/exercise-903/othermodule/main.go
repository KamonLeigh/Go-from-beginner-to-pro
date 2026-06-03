package main

import (
	"fmt"

	"github.com/sicoyle/printer"
)

func main() {
	msg := printer.PrinterNewUUID()
	fmt.Println(msg)
}
