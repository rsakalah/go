package main

import "fmt"

func main() {
	var ConferenceName = "GO Conference"
	const ConferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("welcome to", ConferenceName, "booking application")
	fmt.Println("we have total of", ConferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

}
