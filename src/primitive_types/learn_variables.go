package main

import "fmt"

func main() {
	// working with variables:

	// single variable definition:
	var myInt int = 42
	fmt.Println("the answer: ", myInt)

	// mulitiple varialbes inferring the types of the variables while initialising:
	var val, ok = "yes", true
	fmt.Println("val is: ", val)
	fmt.Println("ok is: ", ok)

	// variables that are declared must be used!!!
	// when compiling (or doing a go run an error is thrown in case a variable is declared and not used.
	// Example:
	var unusedVariable = "foo"

	// it will not compile in case the following line is commented:
	fmt.Println(unusedVariable)

	// in case a variable shall be declared but ignored, the underscore _ can be used.
	// for instance if a function returns two values (a returned value and an error) and the second varialbe shall be ignored, the definition could be as follows:

	var myVariable, _ = "relevant", true
	// this will compile, even if the _ is not used:
	fmt.Println("myVariable: ", myVariable)

	// varialbes shorthand syntax:
	// the declaration: var myInt int = 16 is equal to:
	// the declaration: myInt := 16
	secondInt := 16
	fmt.Println("variable with shorthand declaration: ", secondInt)

	// variable definition separated from assignment:
	var name string
	name = "gerd"
	fmt.Println("name: ", name)

	// understanding the concept of default values for primitive types in go is important
	// because a primitive type cannot be assigned with nil
	// example:
	var defaultInt int
	fmt.Println("default value of an integer varialbe: ", defaultInt)
	var defaultFloat float64
	fmt.Println("default value of an float64 varialbe: ", defaultFloat)
	var defaultString string
	fmt.Println("default value of a string variable: ", defaultString)
}
