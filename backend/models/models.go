package models

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Hub manages all active rooms and users
type Hub struct {
	Rooms      map[string]*Room
	Register   chan *User
	Unregister chan *User
	Lock       sync.RWMutex
	// TODO: decide if we need broadcast channel for web socket events
}

// Room represents a collaborative coding pad room
type Room struct {
	ID       string
	Name     string
	Owner    string
	Users    map[string]*User
	LastUsed time.Time
	Lock     sync.Mutex
}

func (room *Room) UpdateLastUsed() {
	room.Lock.Lock()
	defer room.Lock.Unlock()
	room.LastUsed = time.Now()
}

func (room *Room) AddUser(user *User) {
	room.Lock.Lock()
	defer room.Lock.Unlock()
	room.Users[user.ID] = user
}

func (room *Room) RemoveUser(userID string) {
	room.Lock.Lock()
	defer room.Lock.Unlock()
	delete(room.Users, userID)
}

func (room *Room) GetUser(userID string) (*User, bool) {
	room.Lock.Lock()
	defer room.Lock.Unlock()
	user, exists := room.Users[userID]
	return user, exists
}

// User represents a connected user in the system
type User struct {
	ID          string
	FirstName   string
	LastName    string
	Email       string
	AccessLevel int
	Conn        *websocket.Conn
	Room        *Room
	Lock        sync.Mutex
	// TODO: decide if we need to consider cursor position
}

func (user *User) SetRoom(room *Room) {
	user.Lock.Lock()
	defer user.Lock.Unlock()
	user.Room = room
}

func (user *User) GetRoom() *Room {
	user.Lock.Lock()
	defer user.Lock.Unlock()
	return user.Room
}
