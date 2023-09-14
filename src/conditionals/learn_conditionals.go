package main

import "fmt"

func main() {
	// in go, an array that does not have a numerical index is called a map.
	// comparable to other programming languages where this is called "associative array" or "dictionary" a map can be created with key/value pairs:

	ages := map[string]int{}
	
	ages["Gerd"] = 56
	ages["Tim"] = 6
	ages["Karolina"] = 8

	fmt.Println(ages)

	// access to a specfic key:
	fmt.Println(ages["Gerd"])


	// conditionals example with if, else if and else
	if ages["Gerd"] < 18 {
		fmt.Println("you cant vote")
	} else if ages["Gerd"] < 67 {
		fmt.Println("not ready for retirement")
	} else	{
		fmt.Println("go to retirement")
	}

	// example with the switch case statement: 
	switch {
		case ages["Gerd"] < 18:
			fmt.Println("you cant vote")
		case ages["Gerd"] < 67:
			fmt.Println("not ready for retirement")
		default:
			fmt.Println("go to retirement")
	}

	// advanced switch statement:
	switch ages["Gerd"] {
		case 1,2,3,5,7,11,13,17,19:
			fmt.Println("age is a small prime number")
		case 16:
			fmt.Println("can drive")
		case 18:
			fmt.Println("can vote")
		case 67:
			fmt.Println("can retire now")
		default:
			fmt.Println("there is noting special about the age")
	}
}
