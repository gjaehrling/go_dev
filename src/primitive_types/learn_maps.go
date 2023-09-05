package main

import "fmt"

func main() {
	// in go, an array that does not have a numerical index is called a map.
	// comparable to other programming languages where this is called "associative array" or "dictionary" a map can be created with key/value pairs:

	// the declaration of a map needs the type of the key in square brakets and the type of the value behind the square brakets
	birthdays := map[string]string{
		"Gerd":     "30.06.1967",
		"Karolina": "28.01.2015",
		"Tim":      "06.06.2017",
	}
	fmt.Println(birthdays)

	ages := map[string]int{}

	ages["Gerd"] = 56
	ages["Tim"] = 6
	ages["Karolina"] = 8

	fmt.Println(ages)

	// access to a specfic key:
	fmt.Println(ages["Gerd"])

	// delete values from a map, giving the map and the key to be deleted:
	delete(birthdays, "Gerd")
	fmt.Println(birthdays)

}
