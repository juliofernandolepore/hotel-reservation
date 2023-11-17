package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/juliofernandolepore/hotel-reservation/db"
	"github.com/juliofernandolepore/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	var params types.UpdateUserParams
	userID := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	err = c.BodyParser(&params)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid} // bson.M is a map
	err = h.userStore.UpdateUser(c.Context(), filter, params)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": userID})
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
	err := h.userStore.UpdateUser(c.Context(), userID)
	if err != nil {
		return err
	}
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
