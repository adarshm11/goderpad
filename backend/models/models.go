package models

import (
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
}

type Room struct {
	ID       string
	Name     string
	Owner    string
	Users    map[string]*User
	LastUsed time.Time
	// when calculating how long ago: use time.Now().Sub(LastUsed).Seconds() and compare to 604800 (one week)
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
}

type Event struct {
	EventType string
	RoomID    string
}
