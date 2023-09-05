package main

import "fmt"

func main() {
	fmt.Println("Simple string\n") //interpreted string literal
	fmt.Println(`
		this is a multi line \n
		statement...
	`)
	fmt.Println("\u2272")
	// go does not allow to use double quotes and single quotes interchangablely for string
	// this is because single quotes are used for "runes". A rune is a single character that could be used in a string
	fmt.Println('G') // this is an example of a rune. It will return the corresponding number of the character
}
