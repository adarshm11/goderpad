package models

import (
	"sync"

	"goderpad/utils"
)

type Hub struct {
	Rooms map[string]*Room
	mu    sync.RWMutex
}

var (
	hub  *Hub
	once sync.Once
)

func GetHub() *Hub {
	once.Do(func() {
		hub = &Hub{
			Rooms: make(map[string]*Room),
		}
	})
	return hub
}

// Thread-safe helper methods
func (h *Hub) GetRoom(roomId string) (*Room, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	room, exists := h.Rooms[roomId]
	return room, exists
}

func (h *Hub) SetRoom(roomId string, room *Room) {
	h.mu.Lock()
	h.Rooms[roomId] = room
	h.mu.Unlock()
}

func (h *Hub) ExpireRooms(maxAge int64) int {
	h.mu.Lock()
	defer h.mu.Unlock()
	now := utils.GetCurrentTimestamp()
	expiredCount := 0

	for roomId, room := range h.Rooms {
		if now-room.CreatedAt > maxAge {
			delete(h.Rooms, roomId)
			expiredCount++
		}
	}
	return expiredCount
}
