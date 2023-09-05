package main

import "fmt"

func main() {
	fmt.Println("Greater than: ", 1 > 2)
	fmt.Println("Less than: ", 1 < 2)
	fmt.Println("Greater or equal than: ", 1 >= 2)

	// even tough that float and integer are two different types of primitives, the equivalent works?!?!
	fmt.Println("Equivalent: ", 4.0 == 4)
	fmt.Println("Not equivalent: ", 4.0 != 4)

	var err error = nil // Initializing with nil

	// Checking if err is not nil (indicating an error)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("No error, everything is fine.")
	}
}
