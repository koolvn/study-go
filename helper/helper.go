package helper

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func ValidateUserInput(firstName string, lastName string, email string) (bool, bool, bool) {
	isFirstNameValid := len(firstName) >= 2
	if !isFirstNameValid {
		log.Println("First name can't contain only 1 letter")
	}
	isLastNameValid := len(lastName) >= 1
	if !isLastNameValid {
		log.Println("Last name can't contain only 1 letter")
	}
	isEmailValid := strings.Contains(email, "@")
	if !isEmailValid {
		log.Println("Wrong email")
	}
	return isFirstNameValid, isLastNameValid, isEmailValid
}

func SendTickets(user UserData) string {
	time.Sleep(5 * time.Second)
	msg := fmt.Sprintf(
		"%v tickets sent to %v %v sent to %v\n",
		user.NumBookedTickets, user.FirstName, user.LastName, user.Email)
	log.Print(msg)
	return msg
}
