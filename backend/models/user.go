package models

import (
	"sync"
)

type User struct {
	ID   string
	Name string
	Room *Room
	mu   sync.RWMutex
	// TODO: cursor position, color, etc.
}

func (user *User) SetRoom(room *Room) {
	user.mu.Lock()
	user.Room = room
	user.mu.Unlock()
	room.AddUser(user)
}

func (user *User) GetRoom() *Room {
	user.mu.RLock()
	defer user.mu.RUnlock()
	return user.Room
}
