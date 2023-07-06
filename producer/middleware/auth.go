package middleware

import (
	"producer/utils"

	"github.com/gofiber/fiber/v2"
)

func ParseUser(c *fiber.Ctx) error {

	var token string = c.Cookies("accessToken")

	if len(token) < 5 {
		cliam, err := utils.GetDataFromToken(token)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "token invalid")
		}
		c.Locals("user", cliam)
		c.Next()
		return nil
	} else {
		return fiber.NewError(fiber.StatusInternalServerError, "token required")
	}

}
