package main

import "fmt"


func main() {

	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Println("There are", conferenceTickets, "conference tickets available.")
	fmt.Println("There are", remainingTickets, "conference tickets remaining.")

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

	remainingTickets -= userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Println("the whole bookings slice: ", bookings)
	fmt.Println("the first booking slice: ", bookings[0])
	fmt.Printf("thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", lastName, firstName, userTickets, email)

	fmt.Println("There are", remainingTickets, "conference tickets remaining.")
}

