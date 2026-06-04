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
	api.Post("/login", handlers.Login)
	api.Get("/getRoutes", handlers.GetRoutes)
	api.Get("/getStation", handlers.GetStation)
	api.Get("/getTrain/:id", handlers.GetTrain)
	api.Get("/getCarriage/:id", handlers.GetCarriage)
	// ! Админы
	api.Post("/createRoute", handlers.CreateRoute)
	api.Post("/createStation", handlers.CreateStation)
	api.Get("/users", handlers.GetAllUsers)
	api.Get("/users/:id", handlers.GetUserByID)

	log.Fatal(app.Listen(":8080"))
}
