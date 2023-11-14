package api

import (
	"context"
	"log"

	"github.com/Tanmoy095/Hotel_Reservation.git/db"
	"github.com/Tanmoy095/Hotel_Reservation.git/types"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore //db.usersttore is a interface
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")
	user, err := h.userStore.GetUserById(ctx, id)
	if err != nil {
		log.Fatal(err)

	}
	return c.JSON(user)

}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {

	u := types.User{
		FirstName: "Aunmoy",
		LastName:  "Dey",
	}
	return c.JSON(u)

}
