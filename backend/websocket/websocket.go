package websocket

import (
	"goderpad/models"
)

var hub = &models.Hub{
	Rooms:      make(map[string]*models.Room),
	Register:   make(chan *models.Client),
	Unregister: make(chan *models.Client),
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
			Clients: make(map[string]*models.Client),
		}
	}
}

// unregisterRoom removes all clients from the room and deletes the room from the hub
func unregisterRoom(roomId string) {
	if room, roomExists := hub.Rooms[roomId]; roomExists {
		for clientId := range room.Clients {
			client := room.Clients[clientId]
			client.Room = nil
			hub.Unregister <- client
		}
		delete(hub.Rooms, roomId)
	}
}
