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
