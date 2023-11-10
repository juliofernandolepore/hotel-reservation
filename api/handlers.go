package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/juliofernandolepore/hotel-reservation/db"
	"github.com/juliofernandolepore/hotel-reservation/types"
)

type UserHandler struct {
	//using interface
	userStore db.UserStore
}

// constructor
func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func HandlerGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Fer",
		LastName:  "Lepore",
	}
	return c.JSON(u)
}

func (h *UserHandler) HandlerGetUser(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
