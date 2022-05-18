// Package counterparty implements functions for the manipulation of counterparties.
package counterparty

type Employee struct {
	age int
}

type Customer struct {
	age int
}

type Age interface{}

func NewEmployee(age int) *Employee {
	return &Employee{age: age}
}

func NewCustomer(age int) *Customer {
	return &Customer{age: age}
}

func Elder(any ...Age) int {
	var elder int
	for _, a := range any {
		if e, ok := a.(*Employee); ok {
			elder = e.age
		}
		if c, ok := a.(*Customer); ok {
			elder = c.age
		}

	}
	return elder
}
