package main

func main() {
	emailer := &SendEmail{
		From: "hi@welcome.com",
	}

	welcomer := &CustomerWelcome{
		Emailer: emailer,
	}

	welcomer.Welcome("Bob", "bob@smih.com")
}
