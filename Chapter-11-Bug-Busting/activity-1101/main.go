package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	ErrInvalidSSNLength  = errors.New("SSN length is invalid")
	ErrInvalidSSNNumber  = errors.New("SSN non-nurmeric digits")
	ErrInvalidSSNPrefix  = errors.New("SSN incorrect prefix: 000")
	ErrInvalidDigitPlace = errors.New("SSN invalid digit placement")
)

func main() {
	validateSSN := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33- 3333", "087-65-4321", "123-45-zzzz"}
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	log.Printf("Checking data %#v\n", validateSSN)

	for idx, ssn := range validateSSN {
		log.Printf("Validate data %s %d of %d", ssn, idx+1, len(validateSSN))
		ssn = strings.ReplaceAll(ssn, "-", "")

		err := checkLength(ssn)
		if err != nil {
			log.Print(err)
		}

		err = isNumber(ssn)
		if err != nil {
			log.Print(err)
		}

		err = isPrefixValid(ssn)
		if err != nil {
			log.Print(err)
		}

		err = isDigetPlacementValid(ssn)
		if err != nil {
			log.Print(err)
		}
	}

}

func isNumber(ssn string) error {
	_, err := strconv.Atoi(ssn)

	if err != nil {
		return fmt.Errorf("the value %s caused an error %w", ssn, ErrInvalidSSNNumber)
	}

	return nil
}

func isPrefixValid(ssn string) error {
	result := strings.HasPrefix(ssn, "000")

	if result {
		return fmt.Errorf("the value %s caused an error %w", ssn, ErrInvalidSSNPrefix)
	}

	return nil
}

func isDigetPlacementValid(ssn string) error {
	if string(ssn[0]) == "9" && (string(ssn[3]) != "9" && string(ssn[3]) != "7") {
		return fmt.Errorf("the value %s caused an error %w", ssn, ErrInvalidDigitPlace)
	}

	return nil
}

func checkLength(ssn string) error {
	result := strings.TrimSpace(ssn)

	if len(result) != 9 {
		return fmt.Errorf("the value %s caused an error %w", result, ErrInvalidSSNLength)
	}

	return nil
}
