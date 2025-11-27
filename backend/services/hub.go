package services

import (
	"goderpad/models"
)

var hub = &models.Hub{
	Rooms:      make(map[string]*models.Room),
	Register:   make(chan *models.User),
	Unregister: make(chan *models.User),
	Broadcast:  make(chan models.Event),
}

func RegisterUsers() {
	// To-Do: goroutine that registers users to rooms
}

func UnregisterUsers() {
	// To-Do: goroutine that unregisters users from rooms
}

func GetHub() *models.Hub {
	return hub
}
