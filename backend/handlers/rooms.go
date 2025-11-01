package handlers

import (
	"goderpad/models"
)

var hub = &models.Hub{
	Rooms:      make(map[string]*models.Room),
	Register:   make(chan *models.User),
	Unregister: make(chan *models.User),
	Broadcast:  make(chan models.Event),
}

// a lot of these functions will probably need to me moved to different files later

/*
1 Room = websocket
Hub manages websockets
Hub has goroutines (active functions) listening for registration and unregistration of websockets
Goroutines pass information to the Hub via channels
Room has maximum 2 Clients, any successive requests to join are rejected

*/

func registerRoom(roomId string) {
	if _, roomExists := hub.Rooms[roomId]; !roomExists {
		hub.Rooms[roomId] = &models.Room{
			Users: make(map[string]*models.User),
		}
	}
}

// unregisterRoom removes all Users from the room and deletes the room from the hub
func unregisterRoom(roomId string) {
	if room, roomExists := hub.Rooms[roomId]; roomExists {
		for userId := range room.Users {
			user := room.Users[userId]
			user.Room = nil
			hub.Unregister <- user
		}
		delete(hub.Rooms, roomId)
	}
}
