package events

import "reflect"

var Topics = []string{
	reflect.TypeOf(CreateUser{}).Name(),
	reflect.TypeOf(DeleteUser{}).Name(),
	reflect.TypeOf(Login{}).Name(),
}

type Event interface{}

type CreateUser struct {
	UserName string
	Email    string
	Password string
	ID       string
	Role     string
}

type DeleteUser struct {
	ID string
}

type Login struct {
	Email    string
	Password string
}
