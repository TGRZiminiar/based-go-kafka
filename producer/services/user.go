package services

import (
	"events"
	"log"
	"producer/command"
	"producer/repositories"
	"producer/utils"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(command command.CreateUser) (id string, token string, err error)
	Login(command command.Login) (token string, err error)
	CurrentUser(userId string) (token string, err error)
}

type userService struct {
	eventProducer EventProducer
	userRepo      repositories.UserRepository
}

func NewUserService(eventProducer EventProducer, userRepo repositories.UserRepository) UserService {
	return userService{
		eventProducer: eventProducer,
		userRepo:      userRepo,
	}
}

func (us userService) CreateUser(command command.CreateUser) (id string, token string, err error) {

	password, err := utils.HashPassword(command.Password)
	if err != nil {
		return "", "", err
	}

	userId := uuid.NewString()

	event := events.CreateUser{
		UserName: command.UserName,
		Email:    command.Email,
		Password: password,
		ID:       userId,
		Role:     "user",
	}

	token, err = utils.GenToken(utils.CookieCliam{
		UserId:   userId,
		UserName: command.UserName,
		Email:    command.Email,
		Role:     "user",
	})
	if err != nil {
		return "", "", err
	}

	log.Printf("%#v", event)

	return userId, token, us.eventProducer.Produce(event)
}

func (us userService) Login(command command.Login) (token string, err error) {

	user, err := us.userRepo.FindByEmail(command.Email)
	if err != nil {
		return "", err
	}

	err = utils.CheckPassword(user.Password, command.Password)
	if err != nil {
		return "", err
	}

	token, err = utils.GenToken(utils.CookieCliam{
		UserId:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us userService) CurrentUser(userId string) (token string, err error) {

	user, err := us.userRepo.FindIdCurrentUser(userId)
	if err != nil {
		return "", err
	}

	token, err = utils.GenToken(utils.CookieCliam{
		UserId:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}
