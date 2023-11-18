package internal

import "github.com/gofiber/fiber/v2"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{
			"error": err.Error()})
	},
}
