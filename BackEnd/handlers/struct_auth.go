package handlers

import (
	"sellTrainTicket/customType"
	"sync"
)

var (
	nextID = 1
	users  = make(map[int]User)
	mu     sync.RWMutex
)

type User struct {
	ID       int
	DataUser customType.RegisterUser
}
