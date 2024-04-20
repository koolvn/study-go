package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string) (bool, bool, bool) {
	isFirstNameValid := len(firstName) >= 2
	isLastNameValid := len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	return isFirstNameValid, isLastNameValid, isEmailValid
}

// func SendTickets(user UserData)