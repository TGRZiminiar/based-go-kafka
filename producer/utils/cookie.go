package utils

import "github.com/gofiber/fiber/v2"

func SetCookie(name string, token string, c *fiber.Ctx) {

	cookie := fiber.Cookie{
		Name:     name,
		Value:    token,
		HTTPOnly: true,
		SameSite: "Strict",
		Secure:   true,
	}

	c.Cookie(&cookie)

}
