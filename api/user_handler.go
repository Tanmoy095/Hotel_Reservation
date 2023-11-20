package api

import (
	"context"

	"github.com/Tanmoy095/Hotel_Reservation.git/db"
	"github.com/Tanmoy095/Hotel_Reservation.git/types"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore //db.userstore is a interface
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params *types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	user, err := types.NewUserFromParams(*params)
	if err != nil {
		return err

	}
	//now post user to db
	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)

}
func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetAllUsers(c.Context())
	if err != nil {
		return err

	}
	return c.JSON(users)

}
func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	var (
		params types.UpdateUserParams
		userID = c.Params("id")
	)
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	filter := db.Map{"_id": userID}
	if err := h.userStore.UpdateUser(c.Context(), filter, params); err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": userID})
}

func ErrBadRequest() {
	panic("unimplemented")
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")
	user, err := h.userStore.GetUserById(ctx, id)
	if err != nil {
		return err

	}
	return c.JSON(user)

}

func (h *UserHandler) HandleDeleteUser(c fiber.Ctx) error {
	var userID = c.Params("id")
	if err := h.userStore.DeleteUser(c.Context(), userID); err != nil {
		return err
	}
	return c.JSON(map[string]string{"deleted": userID})
}
