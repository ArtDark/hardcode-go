// Package counterparty implements functions for the manipulation of counterparties.
package counterparty

type Employee struct {
	age int
}

type Customer struct {
	age int
}

type Age interface {
	Age() int
}

func NewEmployee(age int) *Employee {
	return &Employee{age: age}
}

func NewCustomer(age int) *Customer {
	return &Customer{age: age}
}

func (e *Employee) Age() int {
	return e.age
}

func (c *Customer) Age() int {
	return c.age
}

func Elder(any ...Age) int {
	var elder int
	for _, a := range any {
		if a.Age() > elder {
			elder = a.Age()
		}
	}
	return elder
}
