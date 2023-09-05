package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Addition: ", 1+3) // the experession 1+3 will be calculated as an experession before it put to the fmt.Println function
	fmt.Println("Substraction: ", 27-13)
	fmt.Println("Multiplication: ", 9*11)
	fmt.Println("Division: ", 20/4)

	// integers vs. floats
	fmt.Println("Important concept in go: types are not going to be converted.\nExample: ")
	fmt.Println("Division of 20 divided by 3 =  ", 20/3)     // when working with Integers, go will give us an Integer back
	fmt.Println("Division of 20.0 divided by 3 =  ", 20.0/3) // when working with Floats, go will give us a Float back

	// use the math functions:
	fmt.Println("Exponents: ", math.Pow(7, 3))
}
