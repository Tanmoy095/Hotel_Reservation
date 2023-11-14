package db

import (
	"context"

	"github.com/Tanmoy095/Hotel_Reservation.git/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const USERCOLLECTION = "users"

type UserStore interface {
	GetUserById(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(USERCOLLECTION),
	}
}

func (s *MongoUserStore) GetUserById(ctx context.Context, id string) (*types.User, error) {
	//validates the correctnes of the id
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err

	}
	var user types.User
	if err := s.coll.FindOne(ctx, bson.M{"_id": objID}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil

}
