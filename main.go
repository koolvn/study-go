package main

import (
	"fmt"
	"log"

	"github.com/koolvn/study-go.git/helper"
)

const conferenceTickets int = 50

var conferenceName string = "Go conference"
var remainingTickets uint = 50
var bookings = make([]helper.UserData, 0)

func greetUsers() {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your name:\n")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your lastname:\n")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email:\n")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets you want to buy:\n")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(user helper.UserData) {
	bookings = append(bookings, user)
	remainingTickets -= user.NumBookedTickets
}

func main() {
	logger := log.Default()
	for {
		greetUsers()
		firstName, lastName, email, userTickets := getUserInputs()
		user := helper.UserData{
			FirstName:        firstName,
			LastName:         lastName,
			Email:            email,
			NumBookedTickets: userTickets,
		}
		if userTickets > remainingTickets {
			log.Printf("You can't book more than %v remaining tickets!\n", remainingTickets)
			continue
		}
		isFirstNameValid, isLastNameValid, isEmailValid := helper.ValidateUserInput(firstName, lastName, email)

		if isFirstNameValid && isLastNameValid && isEmailValid {
			bookTicket(user)
			logger.Printf("User %v booked %v tickets\n", firstName, userTickets)
			logger.Printf("%v out of %v tickets remaining", remainingTickets, conferenceTickets)
			helper.SendTickets(user)
			if remainingTickets == 0 {
				logger.Printf("Sold out!")
				break
			}
		}

	}
}
