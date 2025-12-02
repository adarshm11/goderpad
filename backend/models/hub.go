package models

import (
	"sync"
)

// The Hub manages all active WebSocket connections (Rooms)
// The Register and Unregister channels handle creation and deletion of the WebSocket connections.
type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Room
	Unregister chan *Room
	mu         sync.RWMutex

	// TODO: decide if we need a channel to broadcast messages
}

func (hub *Hub) GetRoom(roomID string) *Room {
	hub.mu.RLock()
	defer hub.mu.RUnlock()
	room, exists := hub.Rooms[roomID]
	if !exists {
		return nil
	}
	return room
}
