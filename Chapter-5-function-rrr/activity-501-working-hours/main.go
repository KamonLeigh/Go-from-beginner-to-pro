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

func main() {
	employee := developer{
		individual: Employee{
			id:        1,
			firstName: "mexico",
			lastName:  "vegas",
		},
		HourlyRate: 10,
	}

	employee.LogHours(Monday, 8)
	employee.LogHours(Tuesday, 10)

	fmt.Println("Hours worked on Monday:", employee.weekday[Monday])
	fmt.Println("Hours worked on Tuesday:", employee.weekday[Tuesday])
	fmt.Println("Hours worked this week", employee.WorkedHours())
}
