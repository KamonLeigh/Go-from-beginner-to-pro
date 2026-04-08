package main

import (
	"fmt"
	"unicode"
)

func passwordChecker(pw string) bool {
	pwR := []rune(pw)

	if len(pwR) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, val := range pwR {
		if unicode.IsUpper(val) {
			hasUpper = true
		}

		if unicode.IsLower(val) {
			hasLower = true
		}

		if unicode.IsNumber(val) {
			hasNumber = true
		}

		if unicode.IsPunct(val) || unicode.IsSymbol(val) {
			hasSymbol = true
		}

	}
	return hasLower && hasNumber && hasSymbol && hasUpper
}

func main() {
	if passwordChecker("") {
		fmt.Println("Password Good")
	} else {
		fmt.Println("Password Bad")
	}

	if passwordChecker("This!I5A") {
		fmt.Println("Password Good")
	} else {
		fmt.Println("Password Bad")
	}
}
