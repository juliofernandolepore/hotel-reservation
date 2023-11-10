package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/juliofernandolepore/hotel-reservation/api"
	"github.com/juliofernandolepore/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbname = "hotel"
var userColl = "users"

func main() {
	// starting .env config "uri"
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	uri := os.Getenv("uri")
	// mongo dsn conection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}
	// declare a context
	ctx := context.Background()
	// connect to db && collection
	coll := client.Database(dbname).Collection(userColl)
	//create a user
	user1 := types.User{
		FirstName: "joaquin",
		LastName:  "fenix",
	}
	//insert a user
	res, err := coll.InsertOne(ctx, user1)
	if err != nil {
		log.Fatalln("error al insertar en la collection ", err)
	}
	//check client info
	log.Println(res)

	//running a server
	port := flag.String("listen in address", ":5000", "API LISTEN ADDRESS")
	flag.Parse()

	app := fiber.New()
	app.Get("/users", api.HandlerGetUsers)
	app.Get("/users/id", api.HandlerGetUser)
	app.Listen(*port)
}
