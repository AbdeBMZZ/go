package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	tickets uint
}

var wg = sync.WaitGroup{}

func main() {



	fmt.Println("There are", conferenceTickets, "conference tickets available.")
	fmt.Println("There are", remainingTickets, "conference tickets remaining.")

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTickets := helper.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTickets {
		bookTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		
		fmt.Println("The first names are: ", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Sorry, there are no more tickets remaining.")
			// break
		}
	} else {

		if !isValidName {
			fmt.Println("first name of last name you entered is invalid is too short") 
		}

		if !isValidEmail {
			fmt.Println("email you entered doesn't contain @ sign")
		}

		if !isValidTickets {
			fmt.Println("number of tickets you entered is invalid")
		}

	}
	
	wg.Wait()
}


func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

	


func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket ( userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	//creating map
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		tickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Println("list of bookings : \n", bookings)

	fmt.Println("the whole bookings slice: ", bookings)
	fmt.Println("the first booking slice: ", bookings[0])
	fmt.Printf("thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", lastName, firstName, userTickets, email)

	fmt.Println("There are", remainingTickets, "conference tickets remaining.")
}


func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, lastName, firstName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")

	wg.Done()
}