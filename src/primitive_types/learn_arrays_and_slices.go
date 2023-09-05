package main

import "fmt"

func main() {
	// creating an array of a fixed length string types including assignment:
	names := [3]string{"Gerd", "Karolina", "Tim"}

	fmt.Println(names)

	// separating declaration from assignment:
	var names2 [4]string
	names2[0] = "Gerd"
	names2[1] = "Karolina"
	names2[2] = "Tim"

	fmt.Println("is names2[3] asigned: ", names2[3] != "")

	var myInt int
	fmt.Println(myInt)

	// growing or schrinking arrays requires slices
	// definition of a slice is comparable to an array, but without defining the length:
	// delaring an empty slice of string values
	names_slice := []string{}

	// for adding values to the slice we use the append function:
	names_slice = append(names_slice, "gerd")

	fmt.Println(names_slice)

	// multiple append:
	names_slice = append(names_slice, "karolina", "tim", "helena")
	fmt.Println(names_slice)

	// the make function can be used to allocate a minimum of values:
	// example:
	names_minimum := make([]string, 4)

	// this avoids that we need to append values all the time
	names_minimum[0] = "gerd"
	names_minimum[1] = "karolina"
	names_minimum[2] = "tim"
	names_minimum[3] = "helena"

	fmt.Println("slice assigned with the make function: ", names_minimum)

	names_minimum = append(names_minimum, "chiara")

	fmt.Println("slice assigned with the make function and appended an additional value: ", names_minimum)

}
