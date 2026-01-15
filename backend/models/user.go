package models

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type User struct {
	UserID         string                `json:"userId"`
	Name           string                `json:"name"`
	CursorPosition CursorPosition        `json:"cursorPosition"`
	Conn           *websocket.Conn       `json:"-"`
	Send           chan BroadcastMessage `json:"-"`
	done           chan struct{}         `json:"-"`
	mu             sync.Mutex            `json:"-"`
}

type CursorPosition struct {
	Line   int `json:"lineNumber"`
	Column int `json:"column"`
}

func CreateUser(userID, name string) *User {
	user := &User{
		UserID:         userID,
		Name:           name,
		CursorPosition: CursorPosition{Line: 1, Column: 1},
		Conn:           nil,
		Send:           make(chan BroadcastMessage),
		done:           make(chan struct{}),
	}
	go user.HandleBroadcasts()
	return user
}

func (u *User) Close() {
	close(u.done)
	close(u.Send)
}

func (u *User) UpdateCursorPosition(line, column int) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.CursorPosition.Line = line
	u.CursorPosition.Column = column
}

// this function reads incoming messages from the Send channel and sends to the user's websocket connection
func (u *User) HandleBroadcasts() {
	for {
		select {
		case <-u.done:
			return
		case msg := <-u.Send:
			if u.Conn != nil {
				err := u.Conn.WriteJSON(msg)
				if err != nil {
					log.Println("Error writing to websocket:", err)
				}
			}
		}
	}
}
