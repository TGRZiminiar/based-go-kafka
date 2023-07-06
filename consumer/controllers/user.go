package controllers

import (
	"consumer/repositories"
	"encoding/json"
	"events"
	repoEvent "events/repositories"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	CreateUser(topic string, eventsBytes []byte) error
}

type userController struct {
	userRepo repositories.UserRepository
}

func NewUserController(userRepo repositories.UserRepository) UserController {
	return userController{userRepo: userRepo}
}

func (repo userController) CreateUser(topic string, eventsBytes []byte) error {
	event := &events.CreateUser{}
	err := json.Unmarshal(eventsBytes, event)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(event.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userData := repoEvent.User{
		ID:       event.ID,
		UserName: event.UserName,
		Email:    event.Email,
		Password: string(hashedPassword),
	}

	err = repo.userRepo.CreateUser(userData)
	if err != nil {
		return err
	}

	log.Printf("[%v] %#v", topic, event)
	return nil
}
