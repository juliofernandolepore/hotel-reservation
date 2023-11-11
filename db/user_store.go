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
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
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

func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	result, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		log.Println("no results in this query")
		return nil, err
	}
	err = result.Decode(&users)
	if err != nil {
		return []*types.User{}, nil
	}
	return users, nil
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
