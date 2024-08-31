package main

import (
	"fmt"
	"rest-api/apiHandler"
	"rest-api/dbConfig"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	fmt.Println("Starting application")
	app := fiber.New(fiber.Config{
		AppName: "Rest-api-sample-Project",
		BodyLimit: 4000 * 1024,
	})

	dbConfig.ConnectToMongoDB()

	apiHandler.Router(app)

	log.Fatal(app.Listen(":3000"))
}