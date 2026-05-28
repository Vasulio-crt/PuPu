package handlers

import (
	"log"
	"sellTrainTicket/myDatabase"

	"github.com/gofiber/fiber/v3"
)

func GetRoutes(c fiber.Ctx) error {
	from := c.Query("from")
	to := c.Query("to")
	date := c.Query("date")

	if from == "" || to == "" || date == "" {
		return c.Status(404).SendString("Missing required query parameters: from, to, date")
	}

	routes, err := myDatabase.GetRoutesDB(from, to, date)
	if err != nil {
		log.Printf("Ошибка при получении маршрутов: %v", err)
		return c.Status(500).SendString("Internal Server Error")
	}
	return c.JSON(routes)
}
