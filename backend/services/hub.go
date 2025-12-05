package services

import (
	"goderpad/models"
)

var hub = &models.Hub{
	Rooms:      make(map[string]*models.Room),
	Register:   make(chan *models.Room),
	Unregister: make(chan *models.Room),
}

func GetHub() *models.Hub {
	return hub
}
