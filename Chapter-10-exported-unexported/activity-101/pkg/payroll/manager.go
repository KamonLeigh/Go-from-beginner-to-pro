package payroll

type Manager struct {
	Person         Employee
	Salary         float64
	CommissionRate float64
}

func (m Manager) Pay() (string, float64) {
	pay := m.Salary + (m.Salary * m.CommissionRate)

	return m.Person.FirstName + " " + m.Person.LastName, pay
}
