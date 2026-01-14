package models

import (
	"github.com/gorilla/websocket"
)

type User struct {
	UserID         string
	Name           string
	CursorPosition CursorPosition
	Conn           *websocket.Conn
	Send           chan BroadcastMessage
}

type CursorPosition struct {
	Line   int
	Column int
}

func CreateUser(userID, name string) *User {
	user := &User{
		UserID:         userID,
		Name:           name,
		CursorPosition: CursorPosition{Line: 0, Column: 0},
		Send:           make(chan BroadcastMessage),
	}
	go user.HandleBroadcasts()
	return user
}

func (u *User) HandleBroadcasts() {
	// this function reads incoming messages from the Send channel
	// then it handles the functionality required and sends to the user's websocket connection
}
