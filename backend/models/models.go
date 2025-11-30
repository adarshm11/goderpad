package models

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Hub manages all active rooms and users
type Hub struct {
	Rooms      map[string]*Room
	Register   chan RegisterRequest
	Unregister chan UnregisterRequest
	Broadcast  chan Event
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
	Deleted  bool
	Lock     sync.RWMutex
}

func (room *Room) UpdateLastUsed() {
	room.Lock.Lock()
	room.LastUsed = time.Now()
	room.Lock.Unlock()
}

func (room *Room) AddUser(user *User) {
	room.Lock.Lock()
	room.Users[user.ID] = user
	room.Lock.Unlock()
}

func (room *Room) RemoveUser(userID string) {
	room.Lock.Lock()
	delete(room.Users, userID)
	room.Lock.Unlock()
}

func (room *Room) GetUser(userID string) (*User, bool) {
	room.Lock.RLock()
	user, exists := room.Users[userID]
	room.Lock.RUnlock()
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
	user.Room = room
	user.Lock.Unlock()
}

func (user *User) GetRoom() *Room {
	user.Lock.Lock()
	defer user.Lock.Unlock()
	return user.Room
}

func (user *User) ClearRoomIfMatch(room *Room) {
	user.Lock.Lock()
	if user.Room == room {
		user.Room = nil
	}
	user.Lock.Unlock()
}

type Event struct {
	EventType string
	RoomID    string
}
