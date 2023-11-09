package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/juliofernandolepore/hotel-reservation/api"
	"github.com/juliofernandolepore/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"
const dbname = "hotel"
const userColl = "users"

func main() {
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
		FirstName: "juan",
		LastName:  "cerutti",
	}
	//insert a user
	coll.InsertOne(ctx, user1)
	//check client info
	log.Println(client)

	//running a server
	port := flag.String("listen in address", ":5000", "API LISTEN ADDRESS")
	flag.Parse()

	app := fiber.New()
	app.Get("/users", api.HandlerGetUsers)
	app.Get("/users/id", api.HandlerGetUser)
	app.Listen(*port)
}
