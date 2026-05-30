package main

import (
	"errors"
	"fmt"
	"os"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Person            Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

type Manager struct {
	Person         Employee
	Salary         float64
	CommissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"

	// d := Developer{Individual: Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	// 	m := Manager{Individual: Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}
	// 	err := d.ReviewRating()

	d := Developer{Person: Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m := Manager{Person: Employee{Id: 2, FirstName: "Mr", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}

	err := d.performanceReview()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	payDetails(d)
	payDetails(m)
}

func (d Developer) Pay() (string, float64) {
	yearPay := d.HoursWorkedInYear * d.HourlyRate

	return d.Person.FirstName + " " + d.Person.LastName, yearPay
}

func (m Manager) Pay() (string, float64) {
	pay := m.Salary + (m.Salary * m.CommissionRate)

	return m.Person.FirstName + " " + m.Person.LastName, pay
}

func (d Developer) performanceReview() error {
	total := 0

	for _, rating := range d.Review {
		rating, err := calculateReviewRating(rating)
		if err != nil {
			return err
		}
		total += rating
	}

	totalReviews := float64(len(d.Review))
	average := float64(total) / totalReviews
	println(average)
	fmt.Printf("Performance review for %s %s: %.2f\n", d.Person.FirstName, d.Person.LastName, average)
	return nil
}

func payDetails(p Payer) {
	fullName, pay := p.Pay()

	fmt.Printf("%s has been paid %.2f for the year\n", fullName, pay)
}

func calculateReviewRating(value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case string:
		rating, err := convertRatingStringtoInt(v)
		if err != nil {
			return 0, err
		}
		return rating, nil
	default:
		return 0, errors.New("unknown type provided")
	}
}

func convertRatingStringtoInt(rating string) (int, error) {

	switch rating {
	case "Excellent":
		return 5, nil
	case "Good":
		return 4, nil
	case "Fair":
		return 3, nil
	case "Poor":
		return 2, nil
	case "Unsatisfactory":
		return 1, nil

	default:
		return 0, errors.New("Invalid ratting provided\n")
	}
}
