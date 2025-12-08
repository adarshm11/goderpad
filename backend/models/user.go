package models

import (
	"github.com/gorilla/websocket"
)

type User struct {
	ID   string
	Name string
	Conn *websocket.Conn
	// TODO: cursor position, color, etc.
}
