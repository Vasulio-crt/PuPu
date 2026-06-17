package handlers

import (
	"sellTrainTicket/customType"
	"sellTrainTicket/myDatabase"

	"github.com/gofiber/fiber/v3"
)

func CreateRoute(c fiber.Ctx) error {
	var req customType.RouteDB
	c.Bind().Body(&req)
	err := myDatabase.CreateRouteDB(req)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).SendString("OK")
}

func CreateStation(c fiber.Ctx) error {
	var req []string
	c.Bind().Body(&req)
	res, err := myDatabase.CreateStationDB(req)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).JSON(res)
}

func CheckingSeats(c fiber.Ctx) error {
	seats := make(map[int][]int, 1)
	c.Bind().Body(&seats)
	if len(seats) == 0 {
		return c.Status(400).SendString("Отсутствуют места")
	}
	ok, err := myDatabase.CheckingSeatsDB(seats)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if !ok {
		return c.Status(409).SendString("Выбранные Места заняты")
	}

	return c.Status(200).SendString("OK")
}

func BookingSeats(c fiber.Ctx) error {
	seats := make(map[int][]int, 1)
	c.Bind().Body(&seats)
	login := c.Get("Login", "")
	if login == "" {
		c.Status(404).SendString("Нет логина")
	}
	if len(seats) == 0 {
		return c.Status(400).SendString("Отсутствуют места")
	}
	if err := myDatabase.BookingSeatsDB(login, seats); err != nil {
		if err.Error() == "User not found" {
			return c.Status(404).SendString(err.Error())
		}
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).SendString("OK")
}
