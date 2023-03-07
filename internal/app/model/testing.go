package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		First_name:  "bot",
		Second_name: "tester",
		Email:       "user@exmaple.org",
		Password:    "password",
	}
}
