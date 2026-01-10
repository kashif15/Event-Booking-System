package auth

import (
	"errors"

	"event-booking-api/internal/common"
	"event-booking-api/internal/user"
)

func Register(name, email, password string) (*user.User, error) {

	hashedpassword, err := common.HashPassword(password)

	if err != nil {
		return nil, err
	}

	newuser := &user.User{
		Name: name,
		Email: email,
		Password: hashedpassword,
		Role: "USER",
	}

	err = user.Create(newuser)

	if err != nil {
		return nil, err
	}

	newuser.Password = ""

	return newuser, err
}

func Login(email, password string) (*user.User, error) {

	existingUser, err := user.GetByEmail(email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = common.CheckPassword(password, existingUser.Password)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	existingUser.Password = ""

	return existingUser, nil
}