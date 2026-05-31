package handlers

import (
	"fmt"
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
		return c.Status(409).SendString("Пользователь с Логином или Почтой уже существует")
	}

	hashedPassword, err := utilities.HashPassword(registerUser.Password)
	if err != nil {
		return c.Status(500).SendString("Error hashing password")
	}
	registerUser.Password = hashedPassword

	myDatabase.AddUserDB(registerUser)
	return c.Status(201).SendString(fmt.Sprintf("%s %s", registerUser.Name, registerUser.Surname))
}

func Login(c fiber.Ctx) error {
	loginUser := customType.UserLogin{}

	if err := c.Bind().Body(&loginUser); err != nil {
		return c.Status(400).SendString("Bad Request")
	}
	loginUser.Login = strings.ToLower(loginUser.Login)

	err := myDatabase.GetUserPasswordDB(&loginUser)
	if err != nil {
		log.Printf("Ошибка при поиске пользователя(%s): %v", loginUser.Login, err)
		return c.Status(401).SendString("Неверный логин или пароль")
	}

	if !utilities.CheckPasswordHash(loginUser.Password, loginUser.PasswordDB) {
		return c.Status(401).SendString("Неверный пароль или логин")
	}

	return c.Status(200).SendString(fmt.Sprintf("%s %s", loginUser.Name, loginUser.Surname))
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
