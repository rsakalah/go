package main

import "fmt"

func main() {

	// example 1
	var i int
	i = 42
	fmt.Println(i)

	var ConferenceName = "GO Conference"
	const ConferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("welcome to", ConferenceName, "booking application")
	fmt.Println("we have total of", ConferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	//example 2
	var m1 int
	m1 = 2
	fmt.Println(m1)

	var (
		a1 = 2
		a2 = 4
	)
	fmt.Println(a1 + a2)

	//example 3
	var b1 int32
	var b2 int64
	fmt.Println(int64(b1) + b2)

	// short hand decleraing varible
	c1 := 2
	c2 := 3
	fmt.Println(c1 + c2)

	//example 5
	var d1 int
	d1 = 42
	fmt.Println("Value=%v, Type=%T\n", d1, d1)
	fmt.Printf("Value=%v, Type=%T\n", d1, d1)

	//example 6
	red := 65.78
	fmt.Println("value=%v, Type=%T\n", red, red)
	fmt.Printf("value=%v, Type=%T\n", red, red)

	// variable scoping and shadowing

	var j float64 = 5.6
	fmt.Printf("value=%v, Type=%T", j, j)

	if true {
		var j string = "abc"
		fmt.Printf("value=%v, Type=%t", j, j)
	}

	//varible Declaration Block syntax

	var (
		x  = 77
		xy = 6
	)

	fmt.Println(x, xy)

}
