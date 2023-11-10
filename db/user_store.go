package db

import (
	"context"
	"log"

	"github.com/juliofernandolepore/hotel-reservation/info"
	"github.com/juliofernandolepore/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

// constructor of a MongoUserStore Connection
func NewUserMongoStore(c *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: c,
		coll:   c.Database(info.DBname).Collection(info.UserColl), //harcoded
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id) // validate correct ID
	if err != nil {
		log.Println("cant convert to primitive Object ID")
		return nil, err
	}
	var user types.User
	err = s.coll.FindOne(ctx, bson.M{"_id": ObjectID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	log.Println("user finded")
	return &user, nil
}

type PostgresUserStore struct{}
