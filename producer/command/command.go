package command

import "errors"

type CreateUser struct {
	UserName string `json:username`
	Password string `json:password`
	Email    string `json:email`
	ID       string `json:id`
}

func (cu CreateUser) Validator(user CreateUser) error {

	if user.UserName == "" {
		return errors.New("username is required")
	} else if user.Email == "" {
		return errors.New("email is required")
	} else if user.Password == "" {
		return errors.New("password is required")
	} else {
		return nil
	}
}

type Login struct {
	Email    string
	Password string
}

func (lo Login) Validator(login Login) error {
	if login.Email == "" {
		return errors.New("email is required")
	} else if login.Password == "" {
		return errors.New("password is required")
	} else {
		return nil
	}
}
