package payroll

import (
	"errors"
	"fmt"
)

type Developer struct {
	Person            Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

func (d Developer) Pay() (string, float64) {
	yearPay := d.HoursWorkedInYear * d.HourlyRate

	return d.Person.FirstName + " " + d.Person.LastName, yearPay
}

func (d Developer) PerformanceReview() error {
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
