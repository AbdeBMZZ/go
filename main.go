package main

import (
	"fmt"
	"strings"
)


func main() {

	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Println("There are", conferenceTickets, "conference tickets available.")
	fmt.Println("There are", remainingTickets, "conference tickets remaining.")

	for{

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		//pointer to firstName
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Println("Sorry, there are only", remainingTickets, "tickets remaining.")
			continue
		}
		remainingTickets -= userTickets
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Println("the whole bookings slice: ", bookings)
		fmt.Println("the first booking slice: ", bookings[0])
		fmt.Printf("thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", lastName, firstName, userTickets, email)

		fmt.Println("There are", remainingTickets, "conference tickets remaining.")

		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)			
			firstNames = append(firstNames, names[0])
		}

		fmt.Println("The first names of our bookins: ", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Sorry, there are no more tickets remaining.")
			break
		}

	}
}

