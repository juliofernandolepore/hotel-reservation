package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juliofernandolepore/hotel-reservation/types"
)

func HandlerGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Fer",
		LastName:  "Lepore",
	}
	return c.JSON(u)
}

func HandlerGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user1": "somebody"})
}
