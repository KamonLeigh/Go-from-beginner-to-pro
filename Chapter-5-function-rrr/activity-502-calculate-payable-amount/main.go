package main

import "fmt"

type Employee struct {
	id        int
	firstName string
	lastName  string
}

type developer struct {
	individual Employee
	HourlyRate int
	weekday    [7]int
}

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

func (d *developer) LogHours(day WeekDay, hours int) {
	d.weekday[day] = hours
}

func (d *developer) WorkedHours() int {

	total := 0

	for _, value := range d.weekday {
		total += value
	}

	return total
}

func nonLoggedHours() func(int) int {
	total := 0

	return func(i int) int {
		total += i
		return total
	}
}

func (d *developer) PayDay() (int, bool) {

	hoursWorked := d.WorkedHours()

	if hoursWorked > 40 {
		overTimeHours := hoursWorked - 40
		pay := (40 * d.HourlyRate) + (d.HourlyRate * 2 * overTimeHours)

		return pay, true

	}

	return hoursWorked * d.HourlyRate, false
}

func (d *developer) PayDetails() {
	for i, hours := range d.weekday {

		switch i {
		case 0:
			fmt.Println("Sunday hours: ", hours)
		case 1:
			fmt.Println("Monday hours: ", hours)
		case 2:
			fmt.Println("Tuesday hours: ", hours)
		case 3:
			fmt.Println("Wednesday hours: ", hours)
		case 4:
			fmt.Println("Thursday hours: ", hours)
		case 5:
			fmt.Println("Friday hours: ", hours)
		case 6:
			fmt.Println("Saturday hours: ", hours)
		}
	}

	fmt.Println("Hours worked this week: ", d.WorkedHours())
	totalSalary, hasOvertime := d.PayDay()
	fmt.Println("Pay for the week: £", totalSalary)
	fmt.Println("Is this overtime pay: ", hasOvertime)
}

func main() {
	employee := developer{
		individual: Employee{
			id:        1,
			firstName: "mexico",
			lastName:  "vegas",
		},
		HourlyRate: 10,
	}

	x := nonLoggedHours()
	fmt.Println("Tracking hours worked thus far today", x(2))
	fmt.Println("Tracking hours worked thus far today", x(3))
	fmt.Println("Tracking hours worked thus far today", x(5))

	employee.LogHours(Monday, 8)
	employee.LogHours(Tuesday, 10)
	employee.LogHours(Wednesday, 10)
	employee.LogHours(Thursday, 10)
	employee.LogHours(Friday, 6)
	employee.LogHours(Saturday, 8)

	employee.PayDetails()

}
