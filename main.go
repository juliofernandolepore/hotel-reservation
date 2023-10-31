package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := flag.String("listen in address", ":5000", "API LISTEN ADDRESS")
	flag.Parse()

	app := fiber.New()

	app.Get("/inicio", HandleHome)
	app.Listen(*port)
}

func HandleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"mensaje": "contenido del mensaje"})
}
