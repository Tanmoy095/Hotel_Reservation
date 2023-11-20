package main

import (
	"context"
	"fmt"

	"log"

	"github.com/Tanmoy095/Hotel_Reservation.git/db"
	"github.com/Tanmoy095/Hotel_Reservation.git/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	ctx        = context.Background()
	hotelStore db.HotelStore
	roomStore  db.RoomStore
)

func seedHotel(name, location string) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
	}

	//with bson.M we are going to specify query
	// handler initialization

	rooms := []types.Room{

		{
			Type:      types.SingleRoomType,
			BasePrice: 99.99,
		},
		{
			Type:      types.DoubleRoomType,
			BasePrice: 99.79,
		},
		{
			Type:      types.DeluxRoomType,
			BasePrice: 499.99,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(context.Background(), &hotel)
	if err != nil {
		log.Fatal(err)

	}
	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(context.Background(), &room)
		if err != nil {
			log.Fatal(err)

		}
		fmt.Println(insertedRoom)

	}

}

func main() {

	seedHotel("Belucia", "France")
	seedHotel("Radisson BLUE", "Chittagong")
	seedHotel("Shelvey", "London")

}
func init() {
	var err error
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)

}
