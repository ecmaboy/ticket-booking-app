package main

import "fmt"

func main() {
	// create variable of type string, name the value as "confName" and assign value "GO Conference"
	var confName string = "GO Conference"
	// create variable of type uint, name the value as "totalTickets" and assign value 100
	var totalTickets uint = 100
	// create variable of type uint, name the value as "availableTickets" and assign value 100
	var availableTickets uint = 100

	fmt.Printf("Welcome to our %v and we have totally %v tickets  and from them %v are available\n", confName, totalTickets, availableTickets)
}
