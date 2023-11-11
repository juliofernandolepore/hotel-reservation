package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/juliofernandolepore/hotel-reservation/api"
	"github.com/juliofernandolepore/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// starting .env config "uri"
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	uri := os.Getenv("uri")

	//running a server
	port := flag.String("listen in address", ":5000", "API LISTEN ADDRESS")
	flag.Parse()
	// mongo dsn conection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}
	//handler initialization - instance with methods
	userHandler := api.NewUserHandler(db.NewUserMongoStore(client))

	app := fiber.New()

	app.Post("/users", userHandler.HandlePostUser)
	app.Get("/users/:id", userHandler.HandlerGetUser)

	app.Listen(*port)
}
