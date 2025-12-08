package models

import (
	"sync"

	"github.com/gorilla/websocket"

	"goderpad/utils"
)

type Room struct {
	ID       string
	RoomName string
	Users    map[string]*User
	Owner    *User
	Conn     *websocket.Conn
	Document string
	LastUsed int64
	mu       sync.RWMutex
}

func (room *Room) AddUser(user *User) {
	room.mu.Lock()
	defer room.mu.Unlock()

	if len(room.Users) == 2 {
		return // Room is full
	}
	room.Users[user.ID] = user
	room.LastUsed = utils.GetCurrentUnixTimestamp()
}

func (room *Room) SetOwner(user *User) {
	room.mu.Lock()
	room.Owner = user
	room.mu.Unlock()
}

func (room *Room) RemoveUser(userID string) bool {
	room.mu.Lock()
	defer room.mu.Unlock()
	delete(room.Users, userID)
	room.LastUsed = utils.GetCurrentUnixTimestamp()
	return len(room.Users) == 0
}

func (room *Room) UpdateLastUsed() {
	room.mu.Lock()
	room.LastUsed = utils.GetCurrentUnixTimestamp()
	room.mu.Unlock()
}
