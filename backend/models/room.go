package models

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Room struct {
	ID       string
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
}

func (room *Room) UpdateLastUsed() {
	room.mu.Lock()
	room.LastUsed = time.Now().Unix()
	room.mu.Unlock()
}
