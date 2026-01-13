package models

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Room struct {
	RoomID      string `json:"roomId"`
	RoomName    string `json:"roomName"`
	Users       []User `json:"users"`
	CreatedAt   int64
	Document    string
	Broadcast   chan BroadcastMessage
	connections map[string]*websocket.Conn
	mu          sync.RWMutex
}

// The following functions handle strictly the contents of the Rooms - all websocket synchronization is handled in the websocket handlers

func (r *Room) AddUser(user User) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Users = append(r.Users, user)
}

func (r *Room) GetUsers() []User {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.Users
}

func (r *Room) RemoveUser(userID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, user := range r.Users {
		if user.UserID == userID {
			r.Users = append(r.Users[:i], r.Users[i+1:]...)
		}
	}
}

func (r *Room) UpdateDocument(content string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Document = content
}

func (r *Room) GetDocument() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.Document
}

func (r *Room) AddConnection(userID string, conn any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	wsConn := conn.(*websocket.Conn)
	r.connections[userID] = wsConn
}

func (r *Room) RemoveConnection(userID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.connections, userID)
}

func (r *Room) StartBroadcaster() {
	go func() {
		for message := range r.Broadcast {
			r.mu.RLock()
			for _, user := range r.Users {
				// Skip the sender
				if user.UserID == message.UserID {
					continue
				}

				// Send to user (non-blocking with timeout)
				select {
				case user.Send <- message:
				case <-time.After(10 * time.Second):
					// Channel full or slow consumer
					log.Printf("Failed to send message to user %s", user.UserID)
				}
			}
			r.mu.RUnlock()
		}
	}()
}
