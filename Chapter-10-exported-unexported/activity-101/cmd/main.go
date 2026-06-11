package main

import (
	"activity-101/pkg/payroll"
	"fmt"
	"os"
)

var employeeReview = make(map[string]interface{})

func init() {
	fmt.Println("Welcome to the Employee and Performnace Review.")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++")

}

func init() {
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"

}

func main() {
	d := payroll.Developer{Person: payroll.Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m := payroll.Manager{Person: payroll.Employee{Id: 2, FirstName: "Mr", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}

	err := d.PerformanceReview()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	payroll.PayDetails(d)
	payroll.PayDetails(m)

}
