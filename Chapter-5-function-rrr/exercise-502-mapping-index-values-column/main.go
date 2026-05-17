package main

import (
	"fmt"
	"strings"
)

func csvHdrCol(header []string) {
	csvHeadersToColmnIndex := make(map[int]string)

	for i, v := range header {
		v = strings.TrimSpace(v)

		switch strings.ToLower(v) {
		case "employee":
			csvHeadersToColmnIndex[i] = v
		case "hours worked":
			csvHeadersToColmnIndex[i] = v

		case "hourlya rate":
			csvHeadersToColmnIndex[i] = v
		}
	}

	fmt.Println(csvHeadersToColmnIndex)
}

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	csvHdrCol(hdr)
	hdr2 := []string{"empid", "employee", "address", "hours worked", "manager", "hourly rate"}
	csvHdrCol(hdr2)
}
