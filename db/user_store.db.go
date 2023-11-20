package db

import (
	"context"

	"github.com/Tanmoy095/Hotel_Reservation.git/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Map map[string]any

const USERCOLLECTION = "users"

type UserStore interface {
	GetUserById(context.Context, string) (*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	GetAllUsers(context.Context) ([]*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

// CreateUser implements UserStore.

func NewMongoUserStore(client *mongo.Client, dbname string) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(dbname).Collection(USERCOLLECTION),
	}
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error {
	oid, err := primitive.ObjectIDFromHex(filter["_id"].(string))
	if err != nil {
		return err
	}
	filter["_id"] = oid
	update := bson.M{"$set": params.ToBSN()}
	_, err = s.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err

	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}
func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err

	}
	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil

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
func (s *MongoUserStore) GetAllUsers(ctx context.Context) ([]*types.User, error) {
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err

	}
	var users []*types.User
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
