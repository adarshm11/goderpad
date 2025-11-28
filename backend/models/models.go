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
	Register   chan *User
	Unregister chan *User
	Broadcast  chan Event
	Lock       sync.RWMutex
}

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
	defer user.Lock.Unlock()
	user.Room = room
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
