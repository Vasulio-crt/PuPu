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

	api := app.Group("/api")

	api.Post("/register", handlers.Register)
	api.Get("/getRoutes", handlers.GetRoutes)
	api.Get("/users", handlers.GetAllUsers)
	api.Get("/users/:id", handlers.GetUserByID)

	log.Fatal(app.Listen(":8080"))
}
