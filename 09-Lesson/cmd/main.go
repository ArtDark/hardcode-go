package main

import (
	count "GoSearch/09-Lesson/pkg"
	"fmt"
)

func main() {
	alice := count.NewEmployee(25)
	bob := count.NewCustomer(35)
	carol := count.NewCustomer(35)
	dan := count.NewEmployee(40)

	fmt.Printf("Самому старшему контрагенту %d лет.", count.Elder(alice, bob, carol, dan))
}
