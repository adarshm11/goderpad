package models

import (
	"log"
	"sync"
)

type Room struct {
	RoomID    string `json:"roomId"`
	RoomName  string `json:"roomName"`
	Users     []User `json:"users"`
	CreatedAt int64
	Document  string
	Broadcast chan BroadcastMessage
	mu        sync.RWMutex
}

func NewRoom(roomID, roomName string, createdAt int64) *Room {
	return &Room{
		RoomID:    roomID,
		RoomName:  roomName,
		Users:     []User{},
		CreatedAt: createdAt,
		Document:  "",
		Broadcast: make(chan BroadcastMessage, 256), // Buffered channel
	}
}

func (r *Room) AddUser(user User) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Users = append(r.Users, user)

	select {
	case r.Broadcast <- BroadcastMessage{
		Type:    "user_joined",
		Payload: user,
	}:
	default:
		// Log dropped message - this indicates a problem
		log.Printf("WARNING: Broadcast channel full, dropped user_joined message for room %s", r.RoomID)
	}
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

			select {
			case r.Broadcast <- BroadcastMessage{
				Type:    "user_left",
				Payload: user,
			}:
			default:
				log.Printf("WARNING: Broadcast channel full, dropped user_left message for room %s", r.RoomID)
			}
			break
		}
	}
}

func (r *Room) UpdateDocument(content string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Document = content

	select {
	case r.Broadcast <- BroadcastMessage{
		Type:    "document_updated",
		Payload: content,
	}:
	default:
		log.Printf("WARNING: Broadcast channel full, dropped document_updated message for room %s", r.RoomID)
	}
}
