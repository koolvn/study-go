package main

import (
	"fmt"
	"github.com/koolvn/study-go.git/helper"
	"log"
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

func bookTicket(userName string, lastName string, email string, numTickets uint) {
	var userData = helper.UserData{
		FirstName:        userName,
		LastName:         lastName,
		Email:            email,
		NumBookedTickets: numTickets,
	}
	bookings = append(bookings, userData)
	remainingTickets -= numTickets
}

func main() {
	logger := log.Default()
	for {
		greetUsers()
		firstName, lastName, email, userTickets := getUserInputs()
		if userTickets > remainingTickets {
			log.Printf("You can't book more than %v remaining tickets!\n", remainingTickets)
			continue
		}
		isFirstNameValid, isLastNameValid, isEmailValid := helper.ValidateUserInput(firstName, lastName, email)
		if !isFirstNameValid {
			logger.Println("First name can't contain only 1 letter")
		}
		if !isLastNameValid {
			logger.Println("Last name can't contain only 1 letter")
		}
		if !isEmailValid {
			logger.Println("Wrong email")
		}

		if isFirstNameValid && isLastNameValid && isEmailValid {
			bookTicket(firstName, lastName, email, userTickets)
			logger.Printf("User %v booked %v tickets\n", firstName, userTickets)
			logger.Printf("%v out of %v tickets remaining", remainingTickets, conferenceTickets)
			if remainingTickets == 0 {
				logger.Printf("Sold out!")
				break
			}
		}

	}
}
