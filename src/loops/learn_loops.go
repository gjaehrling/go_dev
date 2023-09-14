package main

import "fmt"

func main() {
	ages := map[string]int{}

	ages["Gerd"] = 67
	ages["Tim"] = 5
	ages["Karolina"] = 8
	ages["Helena"] = 18

	fmt.Println(ages)

	// for
	for name, age := range ages {
		fmt.Println("name =", name, " is ", age, " years old")
	}

	// iterating of a map using the range operator:
	for name, age := range ages {
		switch age {
		case 1,2,3,5,7,11,13,17,19:
			fmt.Println(name, "'s age is a small prime number")
		case 16:
			fmt.Println(name, " can drive")
		case 18:
			fmt.Println(name, " can vote")
		case 67:
			fmt.Println(name, " can retire now")
		default:
			fmt.Println(fmt.Sprintf("there is noting special about the age of %s", name))
		}
	}

	// access to a specfic key:
	fmt.Println(ages["Gerd"])

	// traditional C style for loop: 
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// another way doing this: 
	a := 0
	for a < 10 {
		fmt.Println("count ", a)
		a++
	}

	// continue and break: 
	a = 0
	for a < 10 {
		if a % 2 == 0 {
			a++
			continue
		} else if a == 5 {
			break 
		}
		fmt.Println(a, "is unequal")
		a++
	}
}
