package main

import "fmt"

type SendEmail struct {
	From string
}

func (sender *SendEmail) Send(to, subject, body string) error {
	fmt.Printf("Send email to %s", to)
	return nil
}
