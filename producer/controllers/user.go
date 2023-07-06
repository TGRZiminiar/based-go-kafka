package controllers

import (
	"producer/command"
	"producer/services"
	"producer/utils"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	CreateUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	CurrentUser(c *fiber.Ctx) error
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return userController{userService}
}

func (uc userController) CreateUser(c *fiber.Ctx) error {

	command := command.CreateUser{}

	if err := c.BodyParser(&command); err != nil {
		// panic(err)
		return c.Status(400).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	err := command.Validator(command)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	id, token, err := uc.userService.CreateUser(command)
	if err != nil {
		// panic(err)
		return c.Status(400).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	utils.SetCookie("accessToken", token, c)

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"msg": "create user success",
		"id":  id,
	})
}

func (uc userController) Login(c *fiber.Ctx) error {
	command := command.Login{}

	if err := c.BodyParser(&command); err != nil {
		return err
	}

	err := command.Validator(command)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": err,
		})
	}

	token, err := uc.userService.Login(command)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": err,
		})
	}

	utils.SetCookie("accessToken", token, c)

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"msg": "login success",
	})

}

func (uc userController) CurrentUser(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(utils.CookieCliam)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"msg": "authorize is require",
		})
	}

	token, err := uc.userService.CurrentUser(user.UserId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": err,
		})
	}

	utils.SetCookie("accessToken", token, c)

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"msg": "login success",
	})
}
