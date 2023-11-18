package main // for somes internal test

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/juliofernandolepore/hotel-reservation/db"
	"github.com/juliofernandolepore/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	ctx        = context.Background()
)

func seedHotel(name string, location string) error {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
	}
	rooms := []types.Room{
		{Type: types.SingleRoom,
			BasePrice: 80.99,
			Price:     99.99},
		{Type: types.DeluxeRoom,
			BasePrice: 800.99,
			Price:     1099.99},
		{Type: types.DobleRoom,
			BasePrice: 500.99,
			Price:     700.99},
	}
	insertedHotel, err := hotelStore.Insert(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		_, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(insertedRoom)

	}
}

func main() {
	err := seedHotel("fernando", "argentina")
	if err != nil {
		log.Fatal(err)
	}

}

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	uri := os.Getenv("uri")
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln("cant concet with uri")
	}

	roomStore := db.NewMongoRoomStore(client, db.HotelStore)
	hotelStore := db.NewMongoHotelStore(client)

}
