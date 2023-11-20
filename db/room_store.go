package db

import (
	"context"

	"github.com/Tanmoy095/Hotel_Reservation.git/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const ROOMCOLLECTION = "rooms"

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection
	HotelStore
}

// CreateUser implements UserStore.

func NewMongoRoomStore(client *mongo.Client, hotelstore HotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		coll:       client.Database(DBNAME).Collection(ROOMCOLLECTION),
		HotelStore: hotelstore,
	}
}
func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	res, err := s.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err

	}
	room.ID = res.InsertedID.(primitive.ObjectID)

	//Update the hotelStore  with this room id
	filter := bson.M{"_id": room.HotelID}
	update := bson.M{"$push": bson.M{"rooms": room.ID}}
	if err := s.HotelStore.UpdateHotelStore(ctx, filter, update); err != nil {
		return nil, err

	}
	return room, nil

}
