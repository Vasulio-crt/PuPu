package main

import (
	"log"
	"sellTrainTicket/handlers"
	"sellTrainTicket/myDatabase"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	myDatabase.Init()
	defer myDatabase.DB.Close()

	app := fiber.New()
	app.Use(cors.New())

	app.Post("api/register", handlers.Register)
	app.Get("/users", handlers.GetAllUsers)
	app.Get("/users/:id", handlers.GetUserByID)

	log.Fatal(app.Listen(":8080"))
}
