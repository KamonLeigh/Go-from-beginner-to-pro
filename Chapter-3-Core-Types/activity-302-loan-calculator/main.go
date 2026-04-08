package main

import (
	"errors"
	"fmt"
)

const (
	good_credit_score = 450
	lowRatio          = 10
	highRatio         = 20
)

var (
	errorDuration    = errors.New("Duration must be multipule of 12")
	errorCreditScore = errors.New("Error with credit score it must be a positive number")
	errorIncome      = errors.New("Loan is unaffordable due to income")
)

func loanCalculator(creditScore int, income float64, loanAmount float64, loanTerm float64) error {

	interestRate := 15.0
	ratio := highRatio
	if creditScore < good_credit_score {
		interestRate = 20.0
		ratio = lowRatio
	}

	if creditScore < 1 {
		return errorCreditScore
	}

	if income < 1 {
		return errorIncome
	}

	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		return errorDuration
	}

	// Calculate interest on total
	//

	totalInterest := loanAmount * (interestRate / 100) * float64((int(loanTerm) / 12))
	totalBalance := totalInterest + loanAmount
	monthlyPayment := totalBalance / loanTerm

	approved := false

	thresholdPayment := income * (float64(ratio) / 100)

	if monthlyPayment < thresholdPayment {
		approved = true
	}

	fmt.Println("Credit Score : ", creditScore)
	fmt.Println("Income : ", income)
	fmt.Println("Loan Amount : ", loanAmount)
	fmt.Println("Loan Term : ", loanTerm)
	fmt.Println("Monthly Payment: ", monthlyPayment)
	fmt.Println("Rate : ", interestRate)
	fmt.Println("Total Cost : ", totalInterest)
	fmt.Println("Approved : ", approved)
	fmt.Println("")

	return nil
}

func main() {
	// approved
	fmt.Println("Applicant 1")
	fmt.Println("------------")
	err := loanCalculator(500, 1000, 1000, 24)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	// denied
	fmt.Println("Applicant 2")
	fmt.Println("------------")
	err = loanCalculator(350, 1000, 10000, 12)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

}
