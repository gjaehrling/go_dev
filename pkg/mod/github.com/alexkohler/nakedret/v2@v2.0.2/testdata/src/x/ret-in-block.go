package x

import "fmt"

func Dummy() (err error) {
	fmt.Println("My Dummy function")
	fmt.Println("This condition only exists to show a nested return")
	if 1 == 1 {
		return // want "naked return in func `Dummy` with 8 lines of code"
	}
	fmt.Println("Greetings")
	return nil
}
