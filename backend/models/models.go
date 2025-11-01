package models

import "github.com/gorilla/websocket"

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
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Event
}

type Room struct {
	ID      string
	Clients map[string]*Client
}

type Client struct {
	ID             string
	Conn           *websocket.Conn
	Room           *Room
	CursorPosition *CursorPosition
}

type Event struct {
	EventType string
	RoomID    string
}

type TextChangeEvent struct {
	Client *Client
	// To-Do: decide how text changes are represented
}

type CursorUpdateEvent struct {
	Client *Client
	OldPos CursorPosition
	NewPos CursorPosition
}

type UserJoinedEvent struct {
	Client *Client
	Room   *Room
}

type UserLeftEvent struct {
	Client *Client
	Room   *Room
}

var eventType = map[string]any{
	EventTextChange:   TextChangeEvent{},
	EventCursorUpdate: CursorUpdateEvent{},
	EventUserJoined:   UserJoinedEvent{},
	EventUserLeft:     UserLeftEvent{},
}
