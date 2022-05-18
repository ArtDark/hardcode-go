package main

import (
	count "GoSearch/09-Lesson/pkg/counterparty"
	icount "GoSearch/09-Lesson/pkg/icounterparty"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Task 1
	a := icount.NewEmployee(25)
	b := icount.NewCustomer(35)
	c := icount.NewCustomer(35)
	d := icount.NewEmployee(40)

	fmt.Printf("Самому старшему контрагенту %d лет.\n", icount.Elder(a, b, c, d))

	// Task 2
	alice := count.NewEmployee(25)
	bob := count.NewCustomer(35)
	carol := count.NewCustomer(35)
	dan := count.NewEmployee(40)

	fmt.Printf("Самому старшему контрагенту %d лет.\n", count.Elder(alice, bob, carol, dan))

	// Task 3
	if err := writeStrings(os.Stdout, "Hello", 7687, -67, 06.546, "World!"); err != nil {
		log.Println(err)
	}

}

func writeStrings(w io.Writer, any ...any) error {
	for _, a := range any {
		if s, ok := a.(string); ok {
			b, err := json.Marshal(s)
			if err != nil {
				return err
			}
			_, err = w.Write(b)
			if err != nil {
				return err
			}
		}

	}
	return nil
}
