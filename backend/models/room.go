package models

import (
	"sync"
	"time"

	"goderpad/utils"
)

type Room struct {
	RoomID    string
	RoomName  string
	CreatedAt time.Time
	Document  string
	Users     map[string]*User
	mu        sync.Mutex
}

func NewRoom(roomID, roomName string) *Room {
	return &Room{
		RoomID:    roomID,
		RoomName:  roomName,
		CreatedAt: time.Now(),
		Document:  utils.DEFAULT_CODE,
		Users:     make(map[string]*User),
	}
}

func (r *Room) AddUser(user *User) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Users[user.UserID] = user
}

func (r *Room) RemoveUser(userID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.Users, userID)
}

func (r *Room) GetCurrentUsers() []User {
	r.mu.Lock()
	defer r.mu.Unlock()
	users := []User{}
	for _, user := range r.Users {
		users = append(users, *user)
	}
	return users
}

func (r *Room) UpdateDocument(content string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Document = content
}
