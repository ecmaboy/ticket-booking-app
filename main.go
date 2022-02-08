package main

import (
	"fmt"
	"strings"
)

func main() {
	var confName string
	var totalTickets uint
	var availableTickets uint
	var booking []string

	confName = "GO Conference"
	totalTickets = 100
	availableTickets = 100

	fmt.Printf("\nWelcome to ur %v\n", confName)
	fmt.Printf("We have totally %v tickets\n", totalTickets)
	fmt.Printf("And from them %v are available for booking\n", availableTickets)

	for len(booking) <= 100 {
		var firstName, lastName, userEmail string
		var userTickets uint

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		var validUserName = len(firstName) >= 6 && len(lastName) >= 6
		if !validUserName {
			fmt.Printf("You've entered invalid information %v %v. Please try again \n", firstName, lastName)
			continue
		}
		fmt.Print("Enter your email: ")
		fmt.Scan(&userEmail)
		var validUserEmail = strings.Contains(userEmail, "@") && strings.Contains(userEmail, ".")
		if !validUserEmail {
			fmt.Printf("You've entered invalid email address: %v Please try again\n", userEmail)
			continue
		}
		fmt.Print("How many tickets do you want to book? ")
		fmt.Scan(&userTickets)
		validUserTickets := availableTickets >= userTickets
		if !validUserTickets {
			fmt.Printf("Right now we don't have %v tickets available. You can book tickets among 1 and %v\n", userTickets, availableTickets)
			continue
		}

		if validUserName && validUserEmail && validUserTickets {
			fmt.Printf("Dear %v %v, thank you for booking %v tickets. You'll get confirmation email to %v\n", firstName, lastName, userTickets, userEmail)
		}

		availableTickets = availableTickets - userTickets
		booking = append(booking, firstName+" "+lastName)

		fmt.Printf("Now we've left %v tickets\n", availableTickets)

		firstNames := []string{} // creating a slice for first names of users
		for _, bookings := range booking {
			var names = strings.Fields(bookings)
			firstNames = append(firstNames, names[0])

		}

		fmt.Printf("Here is the list of bookers' first names: %v\n", firstNames)

		if availableTickets == 0 {
			fmt.Printf("All tickets are sold. Hope to see you next time")
			break
		}
	}
}
