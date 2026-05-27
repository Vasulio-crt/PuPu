package handlers

import "sync"

var (
	nextID = 1
	users  = make(map[int]User)
	mu     sync.RWMutex
)

type RegisterUser struct {
	Login      string `json:"login"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type User struct {
	ID       int
	DataUser RegisterUser
}
