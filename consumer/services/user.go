package services

import (
	"consumer/controllers"
	"events"
	"log"
	"reflect"
)

type EventHandler interface {
	Handle(topic string, eventsBytes []byte)
}

type userEventHandler struct {
	userController controllers.UserController
}

func NewUserEventHandler(userController controllers.UserController) EventHandler {
	return userEventHandler{userController}
}

func (controller userEventHandler) Handle(topic string, eventsBytes []byte) {
	log.Printf("[%v] %#v", topic, eventsBytes)
	switch topic {
	case reflect.TypeOf(events.CreateUser{}).Name():
		err := controller.userController.CreateUser(topic, eventsBytes)
		if err != nil {
			log.Println(err)
		}
	}

	// case reflect.Typeof(events.C

}
