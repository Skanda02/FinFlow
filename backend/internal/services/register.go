package services

import (
	"errors"
)

type RegisterData struct {
	Name string
	Email string
	Password string
}

func RegisterNewUser(data *RegisterData) error {
	return errors.New("Yet to implement")
}
