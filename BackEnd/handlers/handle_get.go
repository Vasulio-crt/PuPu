package handlers

import (
	"log"
	"sellTrainTicket/myDatabase"
	"strconv"
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

func GetTrain(c fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	train, err := myDatabase.GetTrainDB(intId)
	if err != nil {
		log.Printf("Ошибка при получении поездов: %v", err)
		return c.Status(500).SendString("Internal Server Error")
	}
	if len(train) == 0 {
		return c.Status(404).SendString("Состав поезда не найден")
	}
	return c.JSON(train)
}

func GetCarriage(c fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	seats, err := myDatabase.GetCarriageSeatDB(intId)
	if err != nil {
		log.Printf("Ошибка при получении вагонов: %v", err)
		return c.Status(500).SendString("Internal Server Error")
	}
	if len(seats) == 0 {
		return c.Status(404).SendString("Вагон не найден")
	}

	return c.JSON(seats)
}
