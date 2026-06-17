package handlers

import (
	"log"
	"sellTrainTicket/customType"
	"sellTrainTicket/myDatabase"
	"sellTrainTicket/utilities"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func Register(c fiber.Ctx) error {
	registerUser := customType.User{}

	if err := c.Bind().Body(&registerUser); err != nil {
		return c.Status(400).SendString("Bad Request")
	}
	registerUser.Login = strings.ToLower(registerUser.Login)

	exists, err := myDatabase.UserExists(registerUser.Login, registerUser.Email)
	if err != nil {
		log.Printf("Ошибка при проверке существования пользователя: %v", err)
		return c.Status(500).SendString("Internal Server Error")
	}
	if exists {
		return c.Status(409).SendString("Пользователь с Логином или почтой уже существует")
	}

	registerUser.Password, err = utilities.HashPassword(registerUser.Password)
	if err != nil {
		return c.Status(500).SendString("Error hashing password")
	}
	myDatabase.AddUserDB(registerUser)
	registerUser.Password = ""

	return c.Status(201).JSON(registerUser)
}

func Login(c fiber.Ctx) error {
	loginUser := customType.User{}

	if err := c.Bind().Body(&loginUser); err != nil {
		return c.Status(400).SendString("Bad Request")
	}
	loginUser.Login = strings.ToLower(loginUser.Login)

	passwordDB, err := myDatabase.GetLoginUserDB(&loginUser)
	if err != nil {
		log.Printf("Ошибка при поиске пользователя(%s): %v", loginUser.Login, err)
		return c.Status(401).SendString("Неверный логин или пароль")
	}

	if !utilities.CheckPasswordHash(loginUser.Password, passwordDB) {
		return c.Status(401).SendString("Неверный пароль или логин")
	}
	loginUser.Password = ""

	return c.JSON(loginUser)
}

// !-- Для админов --!
func GetAllUsers(c fiber.Ctx) error {
	allUsers, err := myDatabase.GetAllUsers()
	if err != nil {
		log.Printf("Ошибка на GetAllUsers %v", err)
		return c.Status(500).SendString("Internal Server Error")
	}

	return c.JSON(allUsers)
}

func GetUserByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	user, err := myDatabase.GetUser(id)
	if err != nil {
		log.Printf("Ошибка на GetUserByID %v", err)
		return c.Status(404).SendString("User not found")
	}

	return c.JSON(user)
}
