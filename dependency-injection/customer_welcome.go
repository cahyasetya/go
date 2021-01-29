package main

import "fmt"

type EmailSender interface {
	Send(to, subject, body string) error
}

type CustomerWelcome struct {
	Emailer EmailSender
}

func (welcomer *CustomerWelcome) Welcome(name, email string) error {
	body := fmt.Sprintf("Hi, %s!", name)
	subject := "Welcome"

	return welcomer.Emailer.Send(email, subject, body)
}
