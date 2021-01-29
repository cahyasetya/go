package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FakeEmailSender struct {
	mock.Mock
}

func (mock *FakeEmailSender) Send(to, subject, body string) error {
	args := mock.Called(to, subject, body)
	return args.Error(0)
}

func TestCustomerWelcome_Welcome(t *testing.T) {
	emailer := &FakeEmailSender{}
	emailer.On("Send", "bob@smith.com", "Welcome", "Hi, Bob!").Return(nil)

	welcomer := &CustomerWelcome{
		Emailer: emailer,
	}
	err := welcomer.Welcome("Bob", "bob@smith.com")
	assert.NoError(t, err)

	emailer.AssertExpectations(t)
}
