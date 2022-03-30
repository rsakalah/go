package main

import (
	"fmt"
	"strings"
)

func main() {

	// Creating and initializing a
	// Variable with a string
	// Using shorthand declaration
	My_value_1 := "Welcome to Go language"

	// Using var keyword
	var My_value_2 string
	My_value_2 = "Go language"

	//Dislaying strings
	fmt.Println("string 1:", My_value_1)
	fmt.Println("string 2:", My_value_2)

	//Note: string can be empty not null

	//using backticks("")

	//using double-quote
	My_value_3 := "welome to go language"

	//adding escape character
	My_value_4 := "welcome!\ngo language"

	//using backtricks
	My_value_5 := `Hello! go language`

	//Adding escape character
	//in raw literals
	My_value_6 := `Hello!/ngo language`

	//Diplaying strings
	fmt.Println("string 3:", My_value_3)
	fmt.Println("string 4:", My_value_4)
	fmt.Println("string 5:", My_value_5)
	fmt.Println("string 6:", My_value_6)

	//Replacing a word
	My_value_7 := "my name"

	fmt.Println(strings.ReplaceAll(My_value_7, "m", "no"))

}
