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

var stopChan = make(chan struct{})

func RegisterUsers() {
	for {
		select {
		case user := <-hub.Register:
			room := user.Room
			room.Users[user.ID] = user
		case <-stopChan:
			return
		}
	}
}

func UnregisterUsers() {
	for {
		select {
		case user := <-hub.Unregister:
			room := user.Room
			delete(room.Users, user.ID)
			user.Room = nil
		case <-stopChan:
			return
		}
	}
}

func GetHub() *models.Hub {
	return hub
}

func StopHub() {
	close(stopChan)
}
