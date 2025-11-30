package models

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	EventTextChange   = "text-change"
	EventCursorUpdate = "cursor-update"
	EventUserJoined   = "user-joined"
	EventUserLeft     = "user-left"
)

type CursorPosition struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan RegisterRequest
	Unregister chan UnregisterRequest
	Broadcast  chan Event
	Lock       sync.RWMutex
}

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

type User struct {
	ID             string
	FirstName      string
	LastName       string
	Email          string
	AccessLevel    int
	Conn           *websocket.Conn
	Room           *Room
	CursorPosition *CursorPosition
	Lock           sync.Mutex
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

type Event struct {
	EventType string
	RoomID    string
}
