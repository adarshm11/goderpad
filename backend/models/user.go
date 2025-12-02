package models

import (
	"sync"
)

type User struct {
	ID   string
	Name string
	Room *Room
	Mu   sync.RWMutex
	// TODO: cursor position, color, etc.
}

func (user *User) SetRoom(room *Room) {
	user.Mu.Lock()
	user.Room = room
	user.Mu.Unlock()
	room.AddUser(user)
}

func (user *User) GetRoom() *Room {
	user.Mu.RLock()
	defer user.Mu.RUnlock()
	return user.Room
}
