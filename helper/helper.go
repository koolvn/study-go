package helper

import (
	"fmt"
	"log"
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string) (bool, bool, bool) {
	isFirstNameValid := len(firstName) >= 2
	isLastNameValid := len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	if !isFirstNameValid {
		log.Println("First name can't contain only 1 letter")
	}
	if !isLastNameValid {
		log.Println("Last name can't contain only 1 letter")
	}
	if !isEmailValid {
		log.Println("Wrong email")
	}
	return isFirstNameValid, isLastNameValid, isEmailValid
}

func SendTickets(user UserData) string {
	msg := fmt.Sprintf(
		"%v tickets sent to %v %v sent to %v\n",
		user.NumBookedTickets, user.FirstName, user.LastName, user.Email)
	log.Print(msg)
	return msg
}
