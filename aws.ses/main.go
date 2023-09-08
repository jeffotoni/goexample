package main

import (
	"fmt"
	ses "github.com/jeffotoni/gses"
)

func main() {
	MsgHTML := `<h1>Test send Email - SABESP</h1>`
	To := "<seu-email>"
	From := "<email-here>"
	FromMsg := "Message in email Here"
	Titule := "Your Titule Here"

	err := ses.SendEmail(To, From, FromMsg, Titule, MsgHTML)
	fmt.Println(err)
}
