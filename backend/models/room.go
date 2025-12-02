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
	Mu       sync.RWMutex
}

func (room *Room) AddUser(user *User) {
	room.Mu.Lock()
	defer room.Mu.Unlock()

	if len(room.Users) == 2 {
		return // Room is full
	}
	room.Users[user.ID] = user
}

func (room *Room) UpdateLastUsed() {
	room.Mu.Lock()
	room.LastUsed = time.Now().Unix()
	room.Mu.Unlock()
}
