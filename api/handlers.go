package api

import (
	"log"

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

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	err := c.BodyParser(&params)
	if err != nil {
		return err
	}
	err = params.Validate()
	if err != nil {
		return err
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	userInserted, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(userInserted)
}

func (h *UserHandler) HandlerGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(users)
}

func (h *UserHandler) HandlerGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		c.JSON(map[string]string{"message": "not found"})
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandlerUpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	err := h.userStore.UpdateUser()

}
func (h *UserHandler) HandlerDeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	err := h.userStore.DeleteUser(c.Context(), userID)
	if err != nil {
		c.JSON(map[string]string{"the user is not find to delete:": userID})
		return err
	}
	return c.JSON(map[string]string{"deleted": userID}) //feedback client
}
