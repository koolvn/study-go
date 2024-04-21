package helper

import "fmt"

type UserData struct {
	FirstName        string
	LastName         string
	Email            string
	NumBookedTickets uint
}

func (d UserData) ShowData() string{
	msg := fmt.Sprintf("User: %v %v\nEmail: %v\nTickets Booked: %v", d.FirstName, d.LastName, d.Email, d.NumBookedTickets)
	return msg
}
