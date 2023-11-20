package api

import (
	"github.com/Tanmoy095/Hotel_Reservation.git/db"
	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	hotelStore db.HotelStore
	roomStore  db.RoomStore
}

func NewHotelHandler(hotelStore db.HotelStore, roomStore db.RoomStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: hotelStore,
		roomStore:  roomStore,
	}
}

func (s *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	hotels, err := s.hotelStore.GetAllHotels(c.Context())
	if err != nil {
		return err

	}
	return c.JSON(hotels)
}
