package handlers

import (
	"log"
	"sellTrainTicket/myDatabase"
	"time"

	"github.com/gofiber/fiber/v3"
)

func GetRoutes(c fiber.Ctx) error {
	from := c.Query("from")
	to := c.Query("to")
	date := c.Query("date")

	if from == "" || to == "" || date == "" {
		return c.Status(404).SendString("Missing required query parameters: from, to, date")
	}

	now := time.Now()
	if date < now.Format(time.DateOnly) {
		return c.Status(404).SendString("Дата устарела")
	}

	routes, err := myDatabase.GetRoutesDB(from, to, date)
	if err != nil {
		log.Printf("Ошибка при получении маршрутов: %v", err)
		return c.Status(500).SendString("Internal Server Error")
	}

	if len(routes) == 0 {
		return c.Status(404).SendString("Маршрутов не найдено")
	}
	return c.JSON(routes)
}

func GetStation(c fiber.Ctx) error {
	stations, err := myDatabase.GetAllStationDB()
	if err != nil {
		log.Printf("Ошибка при получении станций: %v", err)
		return c.Status(500).SendString("Internal Server Error")
	}
	return c.JSON(stations)
}
