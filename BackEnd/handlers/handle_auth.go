package handlers

import (
	"sellTrainTicket/utilities"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func Register(c fiber.Ctx) error {
	registerUser := RegisterUser{}

	if err := c.Bind().Body(&registerUser); err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	hashedPassword, err := utilities.HashPassword(registerUser.Password)
	if err != nil {
		return c.Status(500).SendString("Error hashing password")
	}

	registerUser.Password = hashedPassword

	mu.Lock()
	user := User{
		ID:       nextID,
		DataUser: registerUser,
	}
	users[nextID] = user
	nextID++
	mu.Unlock()

	return c.Status(201).SendString("User create")
}

func GetAllUsers(c fiber.Ctx) error {
	mu.RLock()
	defer mu.RUnlock()

	allUsers := make([]User, 0, len(users))
	for _, user := range users {
		allUsers = append(allUsers, user)
	}

	return c.JSON(allUsers)
}

func GetUserByID(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	mu.RLock()
	user, exists := users[id]
	mu.RUnlock()

	if !exists {
		return c.Status(404).SendString("User not found")
	}

	return c.JSON(user)
}
