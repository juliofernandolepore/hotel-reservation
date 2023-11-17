package main // for somes internal test

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/juliofernandolepore/hotel-reservation/db"
	"github.com/juliofernandolepore/hotel-reservation/info"
	"github.com/juliofernandolepore/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	uri := os.Getenv("uri")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln("cant concet with uri")
	}

	ctx := context.Background()
	hotelStore := db.NewMongoHotelStore(client, info.DBname)

	hotelItaliano := types.Hotel{
		Name:     "flavia",
		Location: "italy",
	}
	/* room := types.Room{
		Type:      types.SingleRoom,
		BasePrice: 80.99,
		Price:     99.99,
	} */

	insertedResult, err := hotelStore.InsertHotel(ctx, &hotelItaliano)
	if err != nil {
		log.Fatalln("cant insert hotel in db")
	}

	fmt.Println(insertedResult)

}
